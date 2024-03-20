package main

import (
	"VUTTR-API/domain/tool"
	"VUTTR-API/infra/controllers"
	"VUTTR-API/test/mock"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

func TestToolListWhenIsOk(t *testing.T) {

	ctl := controllers.NewToolController(
		mock.NewToolRepositoryMock([]tool.Tool{
			{
				Id:          "123",
				Title:       "Notion",
				Link:        "https://notion.so",
				Description: "All in one tool.",
				Tags: []string{
					"organization",
					"planning",
				},
			},
		}),
	)

	req, err := http.NewRequest("GET", "/tool", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ctl.GetTools())

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"tools":[{"id":"123","title":"Notion","description":"All in one tool.","link":"https://notion.so","tags":["organization","planning"]}]}`

	if strings.TrimSpace(rr.Body.String()) != string(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestToolListWhenHasQuery(t *testing.T) {

	ctl := controllers.NewToolController(
		mock.NewToolRepositoryMock([]tool.Tool{
			{
				Id:          "123",
				Title:       "Notion",
				Link:        "https://notion.so",
				Description: "All in one tool.",
				Tags: []string{
					"organization",
					"planning",
				},
			},
		}),
	)

	req, err := http.NewRequest("GET", "/tool?test=11", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ctl.GetTools())

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"tools":[{"id":"123","title":"Notion","description":"All in one tool.","link":"https://notion.so","tags":["organization","planning"]}]}`

	if strings.TrimSpace(rr.Body.String()) != string(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestToolGetById(t *testing.T) {

	ctl := controllers.NewToolController(
		mock.NewToolRepositoryMock([]tool.Tool{
			{
				Id:          "123",
				Title:       "Notion",
				Link:        "https://notion.so",
				Description: "All in one tool.",
				Tags: []string{
					"organization",
					"planning",
				},
			},
		}),
	)

	rr := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "/tool/123", nil)
	if err != nil {
		t.Fatal(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/tool/{id}", ctl.GetToolById())
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"id":"123","title":"Notion","description":"All in one tool.","link":"https://notion.so","tags":["organization","planning"]}`

	if strings.TrimSpace(rr.Body.String()) != string(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestToolGetByTag(t *testing.T) {

	ctl := controllers.NewToolController(
		mock.NewToolRepositoryMock([]tool.Tool{
			{
				Id:          "123",
				Title:       "Notion",
				Link:        "https://notion.so",
				Description: "All in one tool.",
				Tags: []string{
					"organization",
					"planning",
				},
			},
		}),
	)

	rr := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "/tool?tag=planning", nil)
	if err != nil {
		t.Fatal(err)
	}

	router := mux.NewRouter()
	router.Path("/tool").Queries("tag", "{tag}").HandlerFunc(ctl.GetToolByTag())
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"tools":[{"id":"123","title":"Notion","description":"All in one tool.","link":"https://notion.so","tags":["organization","planning"]}]}`

	if strings.TrimSpace(rr.Body.String()) != string(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestToolGetByTagNotItemTag(t *testing.T) {

	ctl := controllers.NewToolController(
		mock.NewToolRepositoryMock([]tool.Tool{
			{
				Id:          "123",
				Title:       "Notion",
				Link:        "https://notion.so",
				Description: "All in one tool.",
				Tags: []string{
					"organization",
					"planning",
				},
			},
		}),
	)

	rr := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "/tool?tag=teste", nil)
	if err != nil {
		t.Fatal(err)
	}

	router := mux.NewRouter()
	router.Path("/tool").Queries("tag", "{tag}").HandlerFunc(ctl.GetToolByTag())
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"tools":[]}`

	if strings.TrimSpace(rr.Body.String()) != string(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestToolGetByIdNotFound(t *testing.T) {

	ctl := controllers.NewToolController(
		mock.NewToolRepositoryMock([]tool.Tool{
			{
				Id:          "123",
				Title:       "Notion",
				Link:        "https://notion.so",
				Description: "All in one tool.",
				Tags: []string{
					"organization",
					"planning",
				},
			},
		}),
	)

	rr := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "/tool/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/tool/{id}", ctl.GetToolById())
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := ``

	if strings.TrimSpace(rr.Body.String()) != string(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestToolCreate(t *testing.T) {

	ctl := controllers.NewToolController(
		mock.NewToolRepositoryMock([]tool.Tool{}),
	)

	rr := httptest.NewRecorder()

	content := tool.Tool{
		Id:          "123",
		Title:       "Notion",
		Link:        "https://notion.so",
		Description: "All in one tool.",
		Tags: []string{
			"organization",
			"planning",
		},
	}

	jsonData, err := json.Marshal(content)

	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/tool", bytes.NewReader([]byte(jsonData)))
	if err != nil {
		t.Fatal(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/tool", ctl.CreateTool())
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	poolCreated := tool.Tool{}
	json.Unmarshal(rr.Body.Bytes(), &poolCreated)

	if !(strings.Compare(poolCreated.Title, "Notion") == 0 &&
		strings.Compare(poolCreated.Link, "https://notion.so") == 0 &&
		strings.Compare(
			poolCreated.Description,
			"All in one tool.",
		) == 0) {
		t.Errorf("handler returned unexpected body")
	}
}

func TestToolUpdate(t *testing.T) {

	ctl := controllers.NewToolController(
		mock.NewToolRepositoryMock([]tool.Tool{
			{
				Id:          "123",
				Title:       "Notion",
				Link:        "https://notion.so",
				Description: "All in one tool.",
				Tags: []string{
					"organization",
					"planning",
				},
			},
		}),
	)

	rr := httptest.NewRecorder()

	content := tool.Tool{
		Title: "Notion@",
	}

	jsonData, err := json.Marshal(content)

	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("PUT", "/tool/123", bytes.NewReader([]byte(jsonData)))
	if err != nil {
		t.Fatal(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/tool/{id}", ctl.UpdateTool())
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	toolEdited := tool.Tool{}
	json.Unmarshal(rr.Body.Bytes(), &toolEdited)

	if !(strings.Compare(toolEdited.Title, "Notion@") == 0 &&
		strings.Compare(toolEdited.Link, "https://notion.so") == 0 &&
		strings.Compare(
			toolEdited.Description,
			"All in one tool.",
		) == 0) {
		t.Errorf("handler returned unexpected body")
	}
}

func TestToolDelete(t *testing.T) {

	ctl := controllers.NewToolController(
		mock.NewToolRepositoryMock([]tool.Tool{
			{
				Id:          "123",
				Title:       "Notion",
				Link:        "https://notion.so",
				Description: "All in one tool.",
				Tags: []string{
					"organization",
					"planning",
				},
			},
		}),
	)

	rr := httptest.NewRecorder()

	req, err := http.NewRequest("DELETE", "/tool/123", nil)
	if err != nil {
		t.Fatal(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/tool/{id}", ctl.Delete())
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := ``

	if strings.TrimSpace(rr.Body.String()) != string(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
