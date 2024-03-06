package main

import (
	"api/domain/user"
	"api/infra/controllers"
	"api/test/mock"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

func TestUserListWhenIsOk(t *testing.T) {

	ctl := controllers.NewUserController(
		mock.NewUserRepositoryMock([]user.Users{
			{
				Id:       "123",
				Name:     "Amos",
				Email:    "amos@teste.com",
				Age:      "01-01-2001",
				Password: "123",
			},
		}),
	)

	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ctl.GetUsers())

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"users":[{"id":"123","name":"Amos","email":"amos@teste.com","age":"01-01-2001","password":"123"}]}`

	if strings.TrimSpace(rr.Body.String()) != string(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestUserListWhenHasQuery(t *testing.T) {

	ctl := controllers.NewUserController(
		mock.NewUserRepositoryMock([]user.Users{
			*user.NewUser(
				"123",
				"Amos",
				"amos@teste.com",
				"01-01-2001",
				"123",
			),
		}),
	)

	req, err := http.NewRequest("GET", "/users?test=11", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ctl.GetUsers())

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"users":[{"id":"123","name":"Amos","email":"amos@teste.com","age":"01-01-2001","password":"123"}]}`

	if strings.TrimSpace(rr.Body.String()) != string(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestUserGetById(t *testing.T) {

	ctl := controllers.NewUserController(
		mock.NewUserRepositoryMock([]user.Users{
			{
				Id:       "123",
				Name:     "Amos",
				Email:    "amos@teste.com",
				Age:      "01-01-2001",
				Password: "123",
			},
		}),
	)

	rr := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "/users/123", nil)
	if err != nil {
		t.Fatal(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/users/{id}", ctl.GetUserById())
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"id":"123","name":"Amos","email":"amos@teste.com","age":"01-01-2001","password":"123"}`

	if strings.TrimSpace(rr.Body.String()) != string(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestUserGetBy(t *testing.T) {

	ctl := controllers.NewUserController(
		mock.NewUserRepositoryMock([]user.Users{
			{
				Id:       "123",
				Name:     "Amos",
				Email:    "amos@teste.com",
				Age:      "01-01-2001",
				Password: "123",
			},
		}),
	)

	rr := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "/users/123", nil)
	if err != nil {
		t.Fatal(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/users/{id}", ctl.GetUserById())
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"id":"123","name":"Amos","email":"amos@teste.com","age":"01-01-2001","password":"123"}`

	if strings.TrimSpace(rr.Body.String()) != string(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestUserGetByIdNotFound(t *testing.T) {

	ctl := controllers.NewUserController(
		mock.NewUserRepositoryMock([]user.Users{
			{
				Id:       "123",
				Name:     "Amos",
				Email:    "amos@teste.com",
				Age:      "01-01-2001",
				Password: "123",
			},
		}),
	)

	rr := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "/users/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/users/{id}", ctl.GetUserById())
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

func TestUserCreate(t *testing.T) {

	ctl := controllers.NewUserController(
		mock.NewUserRepositoryMock([]user.Users{}),
	)

	rr := httptest.NewRecorder()

	content := user.Users{
		Name:     "Amos",
		Email:    "amos@teste.com",
		Age:      "2020-10-10",
		Password: "123",
	}

	jsonData, err := json.Marshal(content)

	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/users", bytes.NewReader([]byte(jsonData)))
	if err != nil {
		t.Fatal(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/users", ctl.CreateUser())
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	userCreate := user.Users{}
	json.Unmarshal(rr.Body.Bytes(), &userCreate)

	comparePassword := bcrypt.CompareHashAndPassword([]byte(userCreate.Password), []byte("123"))

	if !(strings.Compare(userCreate.Name, "Amos") == 0 &&
		strings.Compare(userCreate.Email, "amos@teste.com") == 0 &&
		strings.Compare(userCreate.Age, "2020-10-10") == 0 &&
		comparePassword == nil) {
		t.Errorf("handler returned unexpected body")
	}
}

func TestUserUpdate(t *testing.T) {

	ctl := controllers.NewUserController(
		mock.NewUserRepositoryMock([]user.Users{
			{
				Id:       "123",
				Name:     "Amos",
				Email:    "amos@teste.com",
				Age:      "2020-10-10",
				Password: "$2a$10$SobAyxJCuCt8eNXMIderX.547C.DmvNcshUUdixGxAfGAjgUcTtN.",
			},
		}),
	)

	rr := httptest.NewRecorder()

	content := user.Users{
		Name: "Amospk2",
	}

	jsonData, err := json.Marshal(content)

	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("PUT", "/users/123", bytes.NewReader([]byte(jsonData)))
	if err != nil {
		t.Fatal(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/users/{id}", ctl.UpdateUser())
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	userEdited := user.Users{}
	json.Unmarshal(rr.Body.Bytes(), &userEdited)

	comparePassword := bcrypt.CompareHashAndPassword([]byte(userEdited.Password), []byte("123"))

	if !(strings.Compare(userEdited.Name, "Amospk2") == 0 &&
		strings.Compare(userEdited.Email, "amos@teste.com") == 0 &&
		strings.Compare(userEdited.Age, "2020-10-10") == 0 &&
		comparePassword == nil) {
		t.Errorf("handler returned unexpected body")
	}
}

func TestUserDelete(t *testing.T) {

	ctl := controllers.NewUserController(
		mock.NewUserRepositoryMock([]user.Users{
			{
				Id:       "123",
				Name:     "Amos",
				Email:    "amos@teste.com",
				Age:      "2020-10-10",
				Password: "123",
			},
		}),
	)

	rr := httptest.NewRecorder()

	req, err := http.NewRequest("DELETE", "/users/123", nil)
	if err != nil {
		t.Fatal(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/users/{id}", ctl.Delete())
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
