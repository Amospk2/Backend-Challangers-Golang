package main

import (
	"devpartner-api/infra/middlewares"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

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
	router.Use(middlewares.ApplicationTypeSet)
	router.HandleFunc("/", dummyHandler)
	router.ServeHTTP(rr, req)

	if !strings.Contains(rr.Header().Get("Content-Type"), "application/json") {
		t.Errorf("handler returned wrong content type: got %v want %v",
			rr.Header().Get("Content-Type"), "application/json")
	}

}
