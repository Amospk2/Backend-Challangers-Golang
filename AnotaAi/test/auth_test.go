package main

import (
	"api/domain/user"
	"api/infra/controllers"
	"api/test/mock"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
)

type LoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func TestLogin(t *testing.T) {

	ctl := controllers.NewAuthController(
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

	content := LoginRequest{
		Email:    "amos@teste.com",
		Password: "123",
	}

	jsonData, err := json.Marshal(content)

	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/login", bytes.NewReader([]byte(jsonData)))
	if err != nil {
		t.Fatal(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/login", ctl.Login())
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	parsedAccessToken, _ := jwt.ParseWithClaims(
		strings.ReplaceAll(rr.Body.String(), `"`, ``), jwt.MapClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET")), nil
		},
	)
	data := parsedAccessToken.Claims.(jwt.MapClaims)["user"].(string)

	if !parsedAccessToken.Valid || strings.Compare(data, "amos@teste.com") == 0 {
		t.Errorf("handler returned unexpected body")
	}
}
