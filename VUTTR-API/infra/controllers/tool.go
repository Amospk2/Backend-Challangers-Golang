package controllers

import (
	"VUTTR-API/domain/tool"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type ToolController struct {
	repository tool.ToolRepository
}

func (c *ToolController) GetTools() http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			tools, err := c.repository.Get()
			if err != nil {
				log.Fatal(err)
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]any{"tools": tools})
		},
	)
}

func (c *ToolController) GetToolById() http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)

			tools, err := c.repository.GetById(vars["id"])

			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				return
			}

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(tools)
		},
	)
}

func (c *ToolController) GetToolByTag() http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)

			tools, err := c.repository.GetByTag(vars["tag"])

			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				return
			}

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]any{"tools": tools})
		},
	)
}

func (c *ToolController) Delete() http.HandlerFunc {
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

func (c *ToolController) UpdateTool() http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			var toolRequest tool.Tool

			tool, err := c.repository.GetById(vars["id"])

			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				return
			}

			if err = json.NewDecoder(r.Body).Decode(&toolRequest); err != nil {
				w.WriteHeader(http.StatusUnprocessableEntity)
				return
			}

			if len(toolRequest.Title) != 0 && toolRequest.Title != "" {
				tool.Title = toolRequest.Title
			}

			if len(toolRequest.Description) != 0 && toolRequest.Description != "" {
				tool.Description = toolRequest.Description
			}

			if len(toolRequest.Link) != 0 && toolRequest.Link != "" {
				tool.Link = toolRequest.Link
			}

			if len(toolRequest.Tags) != 0 {
				tool.Tags = toolRequest.Tags
			}

			if err = c.repository.Update(tool); err != nil {
				fmt.Println(err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(tool)
		},
	)
}

func (c *ToolController) CreateTool() http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			var tool tool.Tool

			err := json.NewDecoder(r.Body).Decode(&tool)

			if err != nil || !tool.Valid() {
				w.WriteHeader(http.StatusUnprocessableEntity)
				return
			}

			findTool, err := c.repository.GetById(tool.Id)

			if err == nil && findTool.Id != "" {
				w.WriteHeader(http.StatusUnprocessableEntity)
				return
			}

			tool.Id = uuid.NewString()

			if err = c.repository.Create(tool); err != nil {
				log.Fatal(err)
			}

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(tool)
		},
	)
}

func NewToolController(r tool.ToolRepository) *ToolController {
	return &ToolController{
		repository: r,
	}
}
