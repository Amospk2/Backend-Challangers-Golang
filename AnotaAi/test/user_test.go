package main

import (
	"api/domain/user"
	"api/infra/controllers"
	"api/test/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUserListWhenIsOk(t *testing.T) {

	ctl := controllers.NewUserController(
		mock.NewUserRepositoryMock([]user.Users{}),
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
	expected := `{"users":[]}`

	if strings.TrimSpace(rr.Body.String()) != string(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestUserListWhenHasQuery(t *testing.T) {

	ctl := controllers.NewUserController(
		mock.NewUserRepositoryMock([]user.Users{
			*user.NewUser(
				"12345",
				"Amós",
				"teste@gmail.com",
				"19",
				"123123",
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
	expected := `{"users":[{"id":"12345","name":"Amós","email":"teste@gmail.com","age":"19","password":"123123"}]}`

	if strings.TrimSpace(rr.Body.String()) != string(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
