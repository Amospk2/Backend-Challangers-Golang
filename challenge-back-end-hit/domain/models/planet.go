package models

type Planet struct {
	Id      string `json:"id"`
	Nome    string `json:"nome,omitempty"`
	Clima   string `json:"clima,omitempty"`
	Terreno string `json:"terreno,omitempty"`
	Filmes  int16  `json:"filmes"`
}

func (planet Planet) Valid() bool {
	if len(planet.Nome) == 0 && planet.Nome == "" {
		return false
	}

	if len(planet.Terreno) == 0 && planet.Terreno == "" {
		return false
	}

	if len(planet.Clima) == 0 && planet.Clima == "" {
		return false
	}

	return true
}
