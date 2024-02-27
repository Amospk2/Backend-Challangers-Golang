package database

import (
	"challenge-back-end-hit/domain/models"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PlanetRepository struct {
	pool *pgxpool.Pool
}

func (db *PlanetRepository) Get() ([]models.Planet, error) {

	planets := make([]models.Planet, 0)

	rows, err := db.pool.Query(context.Background(),
		"select id, nome, terreno, clima, filmes from public.planets",
	)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var planet models.Planet

		err = rows.Scan(
			&planet.Id,
			&planet.Nome,
			&planet.Terreno,
			&planet.Clima,
			&planet.Filmes,
		)

		if err != nil {
			log.Fatal(err)
		}

		planets = append(planets, planet)
	}

	return planets, nil
}

func (db *PlanetRepository) GetById(id string) (models.Planet, error) {

	var planet models.Planet

	err := db.pool.QueryRow(
		context.Background(),
		"select id, nome, terreno, clima, filmes from public.planets where id=$1", id,
	).Scan(
		&planet.Id,
		&planet.Nome,
		&planet.Terreno,
		&planet.Clima,
		&planet.Filmes,
	)

	if err != nil {
		return models.Planet{}, err
	}

	return planet, nil
}

func (db *PlanetRepository) GetByNome(nome string) (models.Planet, error) {

	var planet models.Planet

	err := db.pool.QueryRow(
		context.Background(),
		"select id, nome, terreno, clima, filmes from public.planets where nome=$1", nome,
	).Scan(
		&planet.Id,
		&planet.Nome,
		&planet.Terreno,
		&planet.Clima,
		&planet.Filmes,
	)

	if err != nil {
		return models.Planet{}, err
	}

	return planet, nil
}

func (db *PlanetRepository) Update(data models.Planet) (*models.Planet, error) {
	_, err := db.pool.Exec(
		context.Background(),
		"UPDATE PLANETS SET nome = $1, terreno = $2, clima = $3, filmes= $4 WHERE id = $5",
		data.Nome, data.Terreno, data.Clima, data.Filmes, data.Id,
	)

	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (db *PlanetRepository) Create(data models.Planet) error {
	_, err := db.pool.Exec(
		context.Background(), "INSERT INTO PLANETS VALUES($1,$2,$3, $4, $5)",
		data.Id, data.Nome, data.Terreno, data.Clima, data.Filmes,
	)

	if err != nil {
		return err
	}

	return nil
}

func (db *PlanetRepository) Delete(id string) error {

	_, err := db.pool.Exec(context.Background(), "DELETE FROM PLANETS WHERE id=$1", id)

	if err != nil {
		return err
	}

	return nil
}

func FilmesByName(nome string) int16 {

	// Estruturas para mapear a resposta JSON
	type Planet struct {
		Name           string   `json:"name"`
		RotationPeriod string   `json:"rotation_period"`
		OrbitalPeriod  string   `json:"orbital_period"`
		Diameter       string   `json:"diameter"`
		Climate        string   `json:"climate"`
		Gravity        string   `json:"gravity"`
		Terrain        string   `json:"terrain"`
		SurfaceWater   string   `json:"surface_water"`
		Population     string   `json:"population"`
		Residents      []string `json:"residents"`
		Films          []string `json:"films"`
		Created        string   `json:"created"`
		Edited         string   `json:"edited"`
		URL            string   `json:"url"`
	}

	type ResponseData struct {
		Count    int      `json:"count"`
		Next     string   `json:"next"`
		Previous *string  `json:"previous"` // Usar ponteiro permite valor nulo
		Results  []Planet `json:"results"`
	}

	for page := 1; page <= 6; page++ {
		requestURL := fmt.Sprintf("https://swapi.dev/api/planets/?page=%d&format=json", page)

		res, err := http.Get(requestURL)
		if err != nil {
			fmt.Printf("error making http request: %s\n", err)
			return 0
		}

		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Printf("error making http request: %s\n", err)
			return 0
		}
		var responseData ResponseData
		if err := json.Unmarshal(body, &responseData); err != nil {
			fmt.Println("Erro ao decodificar o JSON:", err)
			return 0
		}

		for _, planet := range responseData.Results {
			fmt.Print(planet.Name == nome)
			if planet.Name == nome {
				return int16(len(planet.Films))
			}
		}

	}

	return 0

}

func NewPlanetRepository(pool *pgxpool.Pool) *PlanetRepository {
	return &PlanetRepository{pool: pool}
}
