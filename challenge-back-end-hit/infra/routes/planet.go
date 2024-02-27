package routes

import (
	"challenge-back-end-hit/infra/controllers"

	"github.com/gorilla/mux"
)

type PlanetRouter struct {
	controller *controllers.PlanetController
}

func (p *PlanetRouter) Load(mux *mux.Router) {
	mux.HandleFunc("/planets", p.controller.GetPlanets()).Methods("GET")
	mux.Path("/planet").HandlerFunc(p.controller.GetPlanetById()).Queries("id", "{id}").Methods("GET")
	mux.HandleFunc("/planet", p.controller.GetPlanetByName()).Queries("nome", "{nome}").Methods("GET")
	mux.HandleFunc("/planet/{id}", p.controller.UpdatePlanet()).Methods("PUT")
	mux.HandleFunc("/planet/{id}", p.controller.Delete()).Methods("DELETE")
	mux.HandleFunc("/planet", p.controller.CreatePlanet()).Methods("POST")
}

func NewPlanetRouter(
	controller *controllers.PlanetController,
) *PlanetRouter {
	return &PlanetRouter{
		controller: controller,
	}
}
