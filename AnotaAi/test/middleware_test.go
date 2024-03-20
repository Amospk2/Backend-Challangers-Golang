package main

import (
	"api/infra/middleware"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
)

func dummyHandler(w http.ResponseWriter, r *http.Request) {}

func TestSetType(t *testing.T) {

	rr := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	router := mux.NewRouter()
	router.Use(middleware.ApplicationTypeSet)
	router.HandleFunc("/", dummyHandler)
	router.ServeHTTP(rr, req)

	if !strings.Contains(rr.Header().Get("Content-Type"), "application/json") {
		t.Errorf("handler returned wrong content type: got %v want %v",
			rr.Header().Get("Content-Type"), "application/json")
	}

}

func TestAuthenticationWithoutToken(t *testing.T) {

	rr := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "/", nil)

	if err != nil {
		t.Fatal(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/", middleware.AuthenticationMiddleware(dummyHandler))
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

}

func TestAuthenticationWithToken(t *testing.T) {

	rr := httptest.NewRecorder()

	accessToken, _ := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user": "123",
			"exp":  time.Now().Add(time.Duration(time.Hour * 1)).Unix(),
		},
	).SignedString([]byte(os.Getenv("SECRET")))

	req, err := http.NewRequest("GET", "/", nil)

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	if err != nil {
		t.Fatal(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/", middleware.AuthenticationMiddleware(dummyHandler))
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
