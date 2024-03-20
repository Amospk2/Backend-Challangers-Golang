package routes

import (
	"VUTTR-API/infra/controllers"
	"VUTTR-API/infra/middleware"

	"github.com/gorilla/mux"
)

type UserRouter struct {
	controller *controllers.UserController
}

func (p *UserRouter) Load(mux *mux.Router) {
	mux.HandleFunc("/user", middleware.AuthenticationMiddleware(p.controller.GetUsers())).Methods("GET")
	mux.HandleFunc("/user/{id}", middleware.AuthenticationMiddleware(p.controller.GetUserById())).Methods("GET")
	mux.HandleFunc("/user/{id}", middleware.AuthenticationMiddleware(p.controller.UpdateUser())).Methods("PUT")
	mux.HandleFunc("/user/{id}", middleware.AuthenticationMiddleware(p.controller.Delete())).Methods("DELETE")
	mux.HandleFunc("/user", p.controller.CreateUser()).Methods("POST")
}

func NewUserRouter(
	controller *controllers.UserController,
) *UserRouter {
	return &UserRouter{
		controller: controller,
	}
}
