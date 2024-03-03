package routes

import (
	"api/infra/controllers"

	"github.com/gorilla/mux"
)

type AuthRouter struct {
	controller *controllers.AuthController
}

func (p *AuthRouter) Load(mux *mux.Router) {
	mux.HandleFunc("/login", p.controller.Login()).Methods("POST")
}

func NewAuthRouter(
	controller *controllers.AuthController,
) *AuthRouter {
	return &AuthRouter{
		controller: controller,
	}
}
