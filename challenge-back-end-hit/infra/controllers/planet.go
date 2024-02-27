package controllers

import (
	"challenge-back-end-hit/domain/models"
	"challenge-back-end-hit/infra/database"
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PlanetController struct {
	repository *database.PlanetRepository
}

func (c *PlanetController) GetPlanets() http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			planets, err := c.repository.Get()
			if err != nil {
				log.Fatal(err)
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]any{"planets": planets})
		},
	)
}

func (c *PlanetController) GetPlanetByName() http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			users, err := c.repository.GetByNome(vars["nome"])

			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				return
			}

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(users)
		},
	)
}

func (c *PlanetController) GetPlanetById() http.HandlerFunc {
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

func (c *PlanetController) Delete() http.HandlerFunc {
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

func (c *PlanetController) UpdatePlanet() http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			var planetRequest models.Planet

			planet, err := c.repository.GetById(vars["id"])

			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				return
			}

			if err = json.NewDecoder(r.Body).Decode(&planetRequest); err != nil {
				w.WriteHeader(http.StatusUnprocessableEntity)
				return
			}

			if len(planetRequest.Nome) != 0 && planetRequest.Nome != "" {
				planet.Nome = planetRequest.Nome
			}

			if len(planetRequest.Terreno) != 0 && planetRequest.Terreno != "" {
				planet.Terreno = planetRequest.Terreno
			}

			if len(planetRequest.Clima) != 0 && planetRequest.Clima != "" {
				planet.Clima = planetRequest.Clima
			}

			planet.Filmes = database.FilmesByName(planet.Nome)

			var nUser *models.Planet
			if nUser, err = c.repository.Update(planet); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(nUser)
		},
	)
}

func (c *PlanetController) CreatePlanet() http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			var planet models.Planet

			err := json.NewDecoder(r.Body).Decode(&planet)

			if err != nil || !planet.Valid() {
				w.WriteHeader(http.StatusUnprocessableEntity)
				return
			}

			findPlanet, err := c.repository.GetByNome(planet.Nome)

			if err == nil && findPlanet.Nome != "" {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			planet.Filmes = database.FilmesByName(planet.Nome)
			planet.Id = uuid.NewString()

			if err = c.repository.Create(planet); err != nil {
				log.Fatal(err)
			}

			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(planet)
		},
	)
}

func NewPlanetController(pool *pgxpool.Pool) *PlanetController {
	return &PlanetController{
		repository: database.NewPlanetRepository(pool),
	}
}
