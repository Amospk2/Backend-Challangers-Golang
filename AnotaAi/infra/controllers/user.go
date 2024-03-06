package controllers

import (
	"api/domain/user"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	repository user.UserRepository
}

func (c *UserController) GetUsers() http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			users, err := c.repository.Get()
			if err != nil {
				log.Fatal(err)
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]any{"users": users})
		},
	)
}

func (c *UserController) GetUserById() http.HandlerFunc {
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

func (c *UserController) Delete() http.HandlerFunc {
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

func (c *UserController) UpdateUser() http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			var userRequest user.Users

			user, err := c.repository.GetById(vars["id"])

			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				return
			}

			if err = json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
				w.WriteHeader(http.StatusUnprocessableEntity)
				return
			}

			if len(userRequest.Name) != 0 && userRequest.Name != "" {
				user.Name = userRequest.Name
			}

			if len(userRequest.Email) != 0 && userRequest.Email != "" {
				user.Email = userRequest.Email
			}

			if len(userRequest.Age) != 0 && userRequest.Age != "" {
				if _, err := time.Parse("2006-01-02", userRequest.Age); err != nil {
					user.Age = userRequest.Age
				}
			}

			if len(userRequest.Password) != 0 && userRequest.Password != "" {
				user.Password = userRequest.Password
			}

			if err = c.repository.Update(user); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(user)
		},
	)
}

func (c *UserController) CreateUser() http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			var user user.Users

			err := json.NewDecoder(r.Body).Decode(&user)

			if err != nil || !user.Valid() {
				w.WriteHeader(http.StatusUnprocessableEntity)
				return
			}

			findUser, err := c.repository.GetByEmail(user.Email)

			if err == nil && findUser.Id != "" {
				w.WriteHeader(http.StatusUnprocessableEntity)
				return
			}

			user.Id = uuid.NewString()
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			user.Password = string(hashedPassword)

			if err = c.repository.Create(user); err != nil {
				log.Fatal(err)
			}

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(user)
		},
	)
}

func NewUserController(r user.UserRepository) *UserController {
	return &UserController{
		repository: r,
	}
}
