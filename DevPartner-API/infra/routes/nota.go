package routes

import (
	"devpartner-api/infra/controllers"

	"github.com/gorilla/mux"
)

type NotaRouter struct {
	controller *controllers.NotaController
}

func (p *NotaRouter) Load(mux *mux.Router) {
	mux.HandleFunc("/nota", p.controller.GetNotas()).Methods("GET")
	mux.HandleFunc("/nota/{id}", p.controller.GetNotaById()).Methods("GET")
	mux.HandleFunc("/nota/{id}", p.controller.UpdateNota()).Methods("PUT")
	mux.HandleFunc("/nota/{id}", p.controller.Delete()).Methods("DELETE")
	mux.HandleFunc("/nota", p.controller.CreateNota()).Methods("POST")
}

func NewNotaRouter(
	controller *controllers.NotaController,
) *NotaRouter {
	return &NotaRouter{
		controller: controller,
	}
}
