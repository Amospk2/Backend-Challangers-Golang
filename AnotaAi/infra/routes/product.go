package routes

import (
	"api/infra/controllers"
	"api/infra/middleware"

	"github.com/gorilla/mux"
)

type ProductRouter struct {
	controller *controllers.ProductController
}

func (p *ProductRouter) Load(mux *mux.Router) {
	mux.HandleFunc("/products", middleware.AuthenticationMiddleware(p.controller.GetProducts())).Methods("GET")
	mux.HandleFunc("/products/{id}", middleware.AuthenticationMiddleware(p.controller.GetProductById())).Methods("GET")
	mux.HandleFunc("/products/{id}", middleware.AuthenticationMiddleware(p.controller.UpdateProduct())).Methods("PUT")
	mux.HandleFunc("/products/{id}", middleware.AuthenticationMiddleware(p.controller.Delete())).Methods("DELETE")
	mux.HandleFunc("/products", middleware.AuthenticationMiddleware(p.controller.CreateProduct())).Methods("POST")
}

func NewProductRouter(
	controller *controllers.ProductController,
) *ProductRouter {
	return &ProductRouter{
		controller: controller,
	}
}
