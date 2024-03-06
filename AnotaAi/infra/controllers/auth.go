package controllers

import (
	"api/domain/user"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AuthController struct {
	repository user.UserRepository
}

func (c *AuthController) Login() http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			var request AuthRequest

			json.NewDecoder(r.Body).Decode(&request)

			user, err := c.repository.GetByEmail(request.Email)

			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))

			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			accessToken, err := jwt.NewWithClaims(
				jwt.SigningMethodHS256,
				jwt.MapClaims{
					"user": user.Id,
					"exp":  time.Now().Add(time.Duration(time.Hour * 1)).Unix(),
				},
			).SignedString([]byte(os.Getenv("SECRET")))

			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			json.NewEncoder(w).Encode(accessToken)
		},
	)
}

func NewAuthController(r user.UserRepository) *AuthController {
	return &AuthController{
		repository: r,
	}
}
