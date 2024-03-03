package routes

import (
	"api/infra/controllers"
	"api/infra/middleware"

	"github.com/gorilla/mux"
)

type CategoryRouter struct {
	controller *controllers.CategoryController
}

func (p *CategoryRouter) Load(mux *mux.Router) {
	mux.HandleFunc("/categorys", middleware.AuthenticationMiddleware(p.controller.GetCategorys())).Methods("GET")
	mux.HandleFunc("/categorys/{id}", middleware.AuthenticationMiddleware(p.controller.GetCategoryById())).Methods("GET")
	mux.HandleFunc("/categorys/{id}", middleware.AuthenticationMiddleware(p.controller.UpdateCategory())).Methods("PUT")
	mux.HandleFunc("/categorys/{id}", middleware.AuthenticationMiddleware(p.controller.Delete())).Methods("DELETE")
	mux.HandleFunc("/categorys", middleware.AuthenticationMiddleware(p.controller.CreateCategory())).Methods("POST")
}

func NewCategoryRouter(
	controller *controllers.CategoryController,
) *CategoryRouter {
	return &CategoryRouter{
		controller: controller,
	}
}
