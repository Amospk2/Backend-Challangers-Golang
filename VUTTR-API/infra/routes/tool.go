package routes

import (
	"VUTTR-API/infra/controllers"
	"VUTTR-API/infra/middleware"

	"github.com/gorilla/mux"
)

type ToolRouter struct {
	controller *controllers.ToolController
}

func (p *ToolRouter) Load(mux *mux.Router) {
	mux.HandleFunc("/tool", middleware.AuthenticationMiddleware(p.controller.GetTools())).Methods("GET")
	mux.HandleFunc("/tool/{id}", middleware.AuthenticationMiddleware(p.controller.GetToolById())).Methods("GET")
	mux.HandleFunc("/tool/{id}", middleware.AuthenticationMiddleware(p.controller.UpdateTool())).Methods("PUT")
	mux.HandleFunc("/tool/{id}", middleware.AuthenticationMiddleware(p.controller.Delete())).Methods("DELETE")
	mux.HandleFunc("/tool", p.controller.CreateTool()).Methods("POST")
}

func NewToolRouter(
	controller *controllers.ToolController,
) *ToolRouter {
	return &ToolRouter{
		controller: controller,
	}
}
