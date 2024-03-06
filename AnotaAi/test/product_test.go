package main

import (
	"api/domain/product"
	"api/infra/controllers"
	"api/test/mock"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
)

func TestProductListWhenIsOk(t *testing.T) {

	ctl := controllers.NewProductController(
		mock.NewProductRepositoryMock([]product.Product{
			{
				Id:          "123",
				Title:       "Arroz",
				Description: "Arroz preto",
				Price:       123,
				Category:    "123",
				OwnerID:     "123",
			},
		}),
		mock.NewSqsServiceMock([]string{}),
	)

	req, err := http.NewRequest("GET", "/product", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ctl.GetProducts())

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"products":[{"id":"123","title":"Arroz","description":"Arroz preto","price":123,"category":"123","ownerID":"123"}]}`

	if strings.TrimSpace(rr.Body.String()) != string(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestProductListWhenHasQuery(t *testing.T) {

	ctl := controllers.NewProductController(
		mock.NewProductRepositoryMock([]product.Product{
			{
				Id:          "123",
				Title:       "Arroz",
				Description: "Arroz preto",
				Price:       123,
				Category:    "123",
				OwnerID:     "123",
			},
		}),
		mock.NewSqsServiceMock([]string{}),
	)

	req, err := http.NewRequest("GET", "/product?test=11", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ctl.GetProducts())

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"products":[{"id":"123","title":"Arroz","description":"Arroz preto","price":123,"category":"123","ownerID":"123"}]}`

	if strings.TrimSpace(rr.Body.String()) != string(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestProductGetById(t *testing.T) {

	ctl := controllers.NewProductController(
		mock.NewProductRepositoryMock([]product.Product{
			{
				Id:          "123",
				Title:       "Arroz",
				Description: "Arroz preto",
				Price:       123,
				Category:    "123",
				OwnerID:     "123",
			},
		}),
		mock.NewSqsServiceMock([]string{}),
	)

	rr := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "/product/123", nil)
	if err != nil {
		t.Fatal(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/product/{id}", ctl.GetProductById())
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"id":"123","title":"Arroz","description":"Arroz preto","price":123,"category":"123","ownerID":"123"}`

	if strings.TrimSpace(rr.Body.String()) != string(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestProductGetBy(t *testing.T) {

	ctl := controllers.NewProductController(
		mock.NewProductRepositoryMock([]product.Product{
			{
				Id:          "123",
				Title:       "Arroz",
				Description: "Arroz preto",
				Price:       123,
				Category:    "123",
				OwnerID:     "123",
			},
		}),
		mock.NewSqsServiceMock([]string{}),
	)

	rr := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "/product/123", nil)
	if err != nil {
		t.Fatal(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/product/{id}", ctl.GetProductById())
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"id":"123","title":"Arroz","description":"Arroz preto","price":123,"category":"123","ownerID":"123"}`

	if strings.TrimSpace(rr.Body.String()) != string(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestProductGetByIdNotFound(t *testing.T) {

	ctl := controllers.NewProductController(
		mock.NewProductRepositoryMock([]product.Product{
			{
				Id:          "123",
				Title:       "Arroz",
				Description: "Arroz preto",
				Price:       123,
				Category:    "123",
				OwnerID:     "123",
			},
		}),
		mock.NewSqsServiceMock([]string{}),
	)

	rr := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "/product/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/product/{id}", ctl.GetProductById())
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := ``

	if strings.TrimSpace(rr.Body.String()) != string(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestProductCreate(t *testing.T) {

	ctl := controllers.NewProductController(
		mock.NewProductRepositoryMock([]product.Product{
			{
				Id:          "123",
				Title:       "Arroz",
				Description: "Arroz preto",
				Price:       123,
				Category:    "123",
				OwnerID:     "123",
			},
		}),
		mock.NewSqsServiceMock([]string{}),
	)

	rr := httptest.NewRecorder()

	content := product.Product{
		Id:          "123",
		Title:       "Arroz",
		Description: "Arroz preto",
		Price:       123,
		Category:    "123",
		OwnerID:     "123",
	}

	jsonData, err := json.Marshal(content)

	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/product", bytes.NewReader([]byte(jsonData)))
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.WithValue(req.Context(), "user", jwt.MapClaims{
		"user": "123",
		"exp":  time.Now().Add(time.Duration(time.Hour * 1)).Unix(),
	})

	req = req.WithContext(ctx)

	router := mux.NewRouter()
	router.HandleFunc("/product", ctl.CreateProduct())
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	productCreate := product.Product{}
	json.Unmarshal(rr.Body.Bytes(), &productCreate)

	fmt.Println(productCreate)
	if !(strings.Compare(productCreate.Title, "Arroz") == 0 &&
		strings.Compare(productCreate.Description, "Arroz preto") == 0 &&
		productCreate.Price == 123 &&
		strings.Compare(productCreate.OwnerID, "123") == 0 &&
		strings.Compare(productCreate.Category, "123") == 0) {
		t.Errorf("handler returned unexpected body")
	}
}

func TestProductUpdate(t *testing.T) {

	ctl := controllers.NewProductController(
		mock.NewProductRepositoryMock([]product.Product{
			{
				Id:          "123",
				Title:       "Arroz",
				Description: "Arroz preto",
				Price:       123,
				Category:    "123",
				OwnerID:     "123",
			},
		}),
		mock.NewSqsServiceMock([]string{}),
	)

	rr := httptest.NewRecorder()

	content := product.Product{
		Title:       "Feijão",
		Description: "Feijão Preto",
	}

	jsonData, err := json.Marshal(content)

	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("PUT", "/product/123", bytes.NewReader([]byte(jsonData)))
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.WithValue(req.Context(), "user", jwt.MapClaims{
		"user": "123",
		"exp":  time.Now().Add(time.Duration(time.Hour * 1)).Unix(),
	})

	req = req.WithContext(ctx)

	router := mux.NewRouter()
	router.HandleFunc("/product/{id}", ctl.UpdateProduct())
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	productEdited := product.Product{}
	json.Unmarshal(rr.Body.Bytes(), &productEdited)

	fmt.Println(productEdited)
	if !(strings.Compare(productEdited.Title, "Feijão") == 0 &&
		strings.Compare(productEdited.Description, "Feijão Preto") == 0 &&
		productEdited.Price == 123 &&
		strings.Compare(productEdited.OwnerID, "123") == 0 &&
		strings.Compare(productEdited.Category, "123") == 0) {
		t.Errorf("handler returned unexpected body")
	}
}

func TestProductDelete(t *testing.T) {

	ctl := controllers.NewProductController(
		mock.NewProductRepositoryMock([]product.Product{
			{
				Id:          "123",
				Title:       "Arroz",
				Description: "Arroz preto",
				Price:       123,
				Category:    "123",
				OwnerID:     "123",
			},
		}),
		mock.NewSqsServiceMock([]string{}),
	)

	rr := httptest.NewRecorder()

	req, err := http.NewRequest("DELETE", "/product/123", nil)
	if err != nil {
		t.Fatal(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/product/{id}", ctl.Delete())
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := ``

	if strings.TrimSpace(rr.Body.String()) != string(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
