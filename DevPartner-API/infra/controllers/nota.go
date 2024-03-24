package controllers

import (
	"devpartner-api/domain/nota"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type NotaController struct {
	repository nota.NotaRepository
}

func (c *NotaController) GetNotas() http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			notas, err := c.repository.Get()
			if err != nil {
				log.Fatal(err)
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]any{"notas": notas})
		},
	)
}

func (c *NotaController) GetNotaById() http.HandlerFunc {
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

func (c *NotaController) Delete() http.HandlerFunc {
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

func (c *NotaController) UpdateNota() http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			var notaRequest nota.Nota

			nota, err := c.repository.GetById(vars["id"])

			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				return
			}

			if err = json.NewDecoder(r.Body).Decode(&notaRequest); err != nil {
				w.WriteHeader(http.StatusUnprocessableEntity)
				return
			}

			if fmt.Sprintf("%T", notaRequest.NumeroNf) == "float32" && notaRequest.NumeroNf > -1 {
				nota.NumeroNf = notaRequest.NumeroNf
			}

			if _, err := time.Parse("2006-01-02", notaRequest.DataNf); err == nil {
				nota.DataNf = notaRequest.DataNf
			}

			if fmt.Sprintf("%T", notaRequest.ValorTotal) == "float32" && notaRequest.ValorTotal > 0 {
				nota.ValorTotal = notaRequest.ValorTotal
			}

			if len(notaRequest.CnpjEmissorNf) > 0 && notaRequest.CnpjEmissorNf != "" {
				nota.CnpjEmissorNf = notaRequest.CnpjEmissorNf
			}

			if len(notaRequest.CnpjDestinatarioNf) > 0 && notaRequest.CnpjDestinatarioNf != "" {
				nota.CnpjDestinatarioNf = notaRequest.CnpjDestinatarioNf
			}

			if err = c.repository.Update(nota); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(nota)
		},
	)
}

func (c *NotaController) CreateNota() http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			var nota nota.Nota

			err := json.NewDecoder(r.Body).Decode(&nota)

			if err != nil || !nota.Valid() {
				w.WriteHeader(http.StatusUnprocessableEntity)
				return
			}

			nota.Id = uuid.NewString()

			if err = c.repository.Create(nota); err != nil {
				log.Fatal(err)
			}

			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(nota)
		},
	)
}

func NewNotaController(
	r nota.NotaRepository,
) *NotaController {
	return &NotaController{
		repository: r,
	}
}
