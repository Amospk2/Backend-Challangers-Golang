package controllers

import (
	"api/domain/category"
	"api/infra/database"
	"api/infra/service"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CategoryController struct {
	repository *database.CategoryRepository
	service    service.SnsService
}

func (c *CategoryController) GetCategorys() http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			categorys, err := c.repository.Get()
			if err != nil {
				log.Fatal(err)
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]any{"categorys": categorys})
		},
	)
}

func (c *CategoryController) GetCategoryById() http.HandlerFunc {
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

func (c *CategoryController) Delete() http.HandlerFunc {
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

func (c *CategoryController) UpdateCategory() http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			var categoryRequest category.Category

			category, err := c.repository.GetById(vars["id"])

			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				return
			}

			if err = json.NewDecoder(r.Body).Decode(&categoryRequest); err != nil {
				w.WriteHeader(http.StatusUnprocessableEntity)
				return
			}

			fmt.Println(11)

			if len(categoryRequest.Title) != 0 && categoryRequest.Title != "" {
				category.Title = categoryRequest.Title
			}

			if len(categoryRequest.Description) != 0 && categoryRequest.Description != "" {
				category.Description = categoryRequest.Description
			}

			if err = c.repository.Update(category); err != nil {
				fmt.Print(err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			contentBody, err := json.Marshal(category)

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			c.service.PublishInTopic(string(contentBody))

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(category)
		},
	)
}

func (c *CategoryController) CreateCategory() http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			var category category.Category

			err := json.NewDecoder(r.Body).Decode(&category)

			if err != nil || !category.Valid() {
				w.WriteHeader(http.StatusUnprocessableEntity)
				return
			}

			findcategory, err := c.repository.GetById(category.Id)

			if err == nil && findcategory.Id != "" {
				w.WriteHeader(http.StatusUnprocessableEntity)
				return
			}

			category.Id = uuid.NewString()
			user := r.Context().Value("user").(map[string]string)
			category.OwnerID = user["user"]

			if err = c.repository.Create(category); err != nil {
				log.Fatal(err)
			}

			contentBody, err := json.Marshal(category)

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			c.service.PublishInTopic(string(contentBody))

			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(category)
		},
	)
}

func NewCategoryController(pool *pgxpool.Pool, s service.SnsService) *CategoryController {
	return &CategoryController{
		repository: database.NewCategoryRepository(pool),
		service:    s,
	}
}
