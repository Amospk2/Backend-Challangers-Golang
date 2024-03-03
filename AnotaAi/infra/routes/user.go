package routes

import (
	"api/infra/controllers"
	"api/infra/middleware"

	"github.com/gorilla/mux"
)

type UserRouter struct {
	controller *controllers.UserController
}

func (p *UserRouter) Load(mux *mux.Router) {
	mux.HandleFunc("/users", middleware.AuthenticationMiddleware(p.controller.GetUsers())).Methods("GET")
	mux.HandleFunc("/users/{id}", middleware.AuthenticationMiddleware(p.controller.GetUserById())).Methods("GET")
	mux.HandleFunc("/users/{id}", middleware.AuthenticationMiddleware(p.controller.UpdateUser())).Methods("PUT")
	mux.HandleFunc("/users/{id}", middleware.AuthenticationMiddleware(p.controller.Delete())).Methods("DELETE")
	mux.HandleFunc("/users", middleware.AuthenticationMiddleware(p.controller.CreateUser())).Methods("POST")
}

func NewUserRouter(
	controller *controllers.UserController,
) *UserRouter {
	return &UserRouter{
		controller: controller,
	}
}
