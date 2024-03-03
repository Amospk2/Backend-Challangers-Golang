package controllers

import (
	"api/domain/product"
	"api/infra/database"
	"api/infra/service"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductController struct {
	repository *database.ProductRepository
	service    service.SnsService
}

func (c *ProductController) GetProducts() http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			products, err := c.repository.Get()
			if err != nil {
				log.Fatal(err)
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]any{"products": products})
		},
	)
}

func (c *ProductController) GetProductById() http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			users, err := c.repository.GetById(vars["id"])

			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				return
			}

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(users)
		},
	)
}

func (c *ProductController) Delete() http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)

			if _, err := c.repository.GetById(vars["id"]); err != nil {
				w.WriteHeader(http.StatusNotFound)
				return
			}

			if err := c.repository.Delete(vars["id"]); err != nil {
				w.WriteHeader(http.StatusNotFound)
				return
			}

			w.WriteHeader(http.StatusOK)
		},
	)
}

func (c *ProductController) UpdateProduct() http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			var productRequest product.Product

			product, err := c.repository.GetById(vars["id"])

			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				return
			}

			if err = json.NewDecoder(r.Body).Decode(&productRequest); err != nil {
				w.WriteHeader(http.StatusUnprocessableEntity)
				return
			}

			if len(productRequest.Title) != 0 && productRequest.Title != "" {
				product.Title = productRequest.Title
			}

			if len(productRequest.Description) != 0 && productRequest.Description != "" {
				product.Description = productRequest.Description
			}

			if len(productRequest.Category) != 0 && productRequest.Category != "" {
				product.Category = productRequest.Category
			}

			if fmt.Sprintf("%T", product.Price) != "int" && product.Price > 0 {
				product.Price = productRequest.Price
			}

			if err = c.repository.Update(product); err != nil {
				fmt.Print(err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			contentBody, err := json.Marshal(product)

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			c.service.PublishInTopic(string(contentBody))

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(product)
		},
	)
}

func (c *ProductController) CreateProduct() http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			var product product.Product

			err := json.NewDecoder(r.Body).Decode(&product)

			if err != nil || !product.Valid() {
				w.WriteHeader(http.StatusUnprocessableEntity)
				return
			}

			findproduct, err := c.repository.GetById(product.Id)

			if err == nil && findproduct.Id != "" {
				w.WriteHeader(http.StatusUnprocessableEntity)
				return
			}

			product.Id = uuid.NewString()
			user := r.Context().Value("user").(jwt.MapClaims)
			product.OwnerID = user["user"].(string)

			if err = c.repository.Create(product); err != nil {
				log.Fatal(err)
			}

			contentBody, err := json.Marshal(product)

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			c.service.PublishInTopic(string(contentBody))

			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(product)
		},
	)
}

func NewProductController(pool *pgxpool.Pool, s service.SnsService) *ProductController {
	return &ProductController{
		repository: database.NewProductRepository(pool),
		service:    s,
	}
}
