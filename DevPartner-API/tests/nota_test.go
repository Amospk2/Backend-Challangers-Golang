package main

import (
	"bytes"
	"devpartner-api/domain/nota"
	"devpartner-api/infra/controllers"
	"devpartner-api/tests/mock"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

func TestNotaListWhenIsOk(t *testing.T) {

	ctl := controllers.NewNotaController(
		mock.NewNotaRepositoryMock([]nota.Nota{
			{
				Id:                 "123",
				NumeroNf:           1.0,
				ValorTotal:         100,
				DataNf:             "2024-01-01",
				CnpjEmissorNf:      "123",
				CnpjDestinatarioNf: "123",
			},
		}),
	)

	req, err := http.NewRequest("GET", "/nota", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ctl.GetNotas())

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"notas":[{"notaFiscalId":"123","numeroNf":1,"valorTotal":100,"dataNf":"2024-01-01","cnpjEmissorNf":"123","cnpjDestinatarioNf":"123"}]}`

	if strings.TrimSpace(rr.Body.String()) != string(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestNotaListWhenHasQuery(t *testing.T) {

	ctl := controllers.NewNotaController(
		mock.NewNotaRepositoryMock([]nota.Nota{
			{
				Id:                 "123",
				NumeroNf:           1.0,
				ValorTotal:         100,
				DataNf:             "2024-01-01",
				CnpjEmissorNf:      "123",
				CnpjDestinatarioNf: "123",
			},
		}),
	)

	req, err := http.NewRequest("GET", "/nota?test=11", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ctl.GetNotas())

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"notas":[{"notaFiscalId":"123","numeroNf":1,"valorTotal":100,"dataNf":"2024-01-01","cnpjEmissorNf":"123","cnpjDestinatarioNf":"123"}]}`

	if strings.TrimSpace(rr.Body.String()) != string(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestNotaGetById(t *testing.T) {

	ctl := controllers.NewNotaController(
		mock.NewNotaRepositoryMock([]nota.Nota{
			{
				Id:                 "123",
				NumeroNf:           1.0,
				ValorTotal:         100,
				DataNf:             "2024-01-01",
				CnpjEmissorNf:      "123",
				CnpjDestinatarioNf: "123",
			},
		}),
	)

	rr := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "/nota/123", nil)
	if err != nil {
		t.Fatal(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/nota/{id}", ctl.GetNotaById())
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"notaFiscalId":"123","numeroNf":1,"valorTotal":100,"dataNf":"2024-01-01","cnpjEmissorNf":"123","cnpjDestinatarioNf":"123"}`

	if strings.TrimSpace(rr.Body.String()) != string(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestNotaGetByIdNotFound(t *testing.T) {

	ctl := controllers.NewNotaController(
		mock.NewNotaRepositoryMock([]nota.Nota{
			{
				Id:                 "123",
				NumeroNf:           1.0,
				ValorTotal:         100,
				DataNf:             "2024-01-01",
				CnpjEmissorNf:      "123",
				CnpjDestinatarioNf: "123",
			},
		}),
	)

	rr := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "/nota/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/nota/{id}", ctl.GetNotaById())
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

func TestNotaCreate(t *testing.T) {

	ctl := controllers.NewNotaController(
		mock.NewNotaRepositoryMock([]nota.Nota{}),
	)

	rr := httptest.NewRecorder()

	content := nota.Nota{
		Id:                 "123",
		NumeroNf:           1.0,
		ValorTotal:         100,
		DataNf:             "2024-01-01",
		CnpjEmissorNf:      "123",
		CnpjDestinatarioNf: "123",
	}

	jsonData, err := json.Marshal(content)

	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/nota", bytes.NewReader([]byte(jsonData)))
	if err != nil {
		t.Fatal(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/nota", ctl.CreateNota())
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	notaCreated := nota.Nota{}
	json.Unmarshal(rr.Body.Bytes(), &notaCreated)

	compare := !(strings.Compare(notaCreated.DataNf, "2024-01-01") == 0 &&
		strings.Compare(notaCreated.CnpjEmissorNf, "123") == 0 &&
		strings.Compare(
			notaCreated.CnpjDestinatarioNf,
			"123",
		) == 0 &&
		notaCreated.NumeroNf == 1.0 &&
		notaCreated.ValorTotal == 100)

	if compare {
		t.Errorf("handler returned unexpected body")
	}
}

func TestNotaUpdate(t *testing.T) {

	ctl := controllers.NewNotaController(
		mock.NewNotaRepositoryMock([]nota.Nota{
			{
				Id:                 "123",
				NumeroNf:           1.0,
				ValorTotal:         100,
				DataNf:             "2024-01-01",
				CnpjEmissorNf:      "123",
				CnpjDestinatarioNf: "123",
			},
		}),
	)

	rr := httptest.NewRecorder()

	content := nota.Nota{
		NumeroNf: 2.0,
	}

	jsonData, err := json.Marshal(content)

	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("PUT", "/nota/123", bytes.NewReader([]byte(jsonData)))
	if err != nil {
		t.Fatal(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/nota/{id}", ctl.UpdateNota())
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	notaEdited := nota.Nota{}
	json.Unmarshal(rr.Body.Bytes(), &notaEdited)

	compare := !(strings.Compare(notaEdited.DataNf, "2024-01-01") == 0 &&
		strings.Compare(notaEdited.CnpjEmissorNf, "123") == 0 &&
		strings.Compare(
			notaEdited.CnpjDestinatarioNf,
			"123",
		) == 0 && notaEdited.NumeroNf == 2 && notaEdited.ValorTotal == 100)

	if compare {
		t.Errorf("handler returned unexpected body")
	}
}

func TestNotaDelete(t *testing.T) {

	ctl := controllers.NewNotaController(
		mock.NewNotaRepositoryMock([]nota.Nota{
			{
				Id:                 "123",
				NumeroNf:           1.0,
				ValorTotal:         100,
				DataNf:             "2024-01-01",
				CnpjEmissorNf:      "123",
				CnpjDestinatarioNf: "123",
			},
		}),
	)

	rr := httptest.NewRecorder()

	req, err := http.NewRequest("DELETE", "/nota/123", nil)
	if err != nil {
		t.Fatal(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/nota/{id}", ctl.Delete())
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
