package nota

import "github.com/google/uuid"

// notaFiscalId, numeroNf, valorTotal, dataNf, cnpjEmissorNf e cnpjDestinatarioNf

type Nota struct {
	Id                 string  `json:"notaFiscalId"`
	NumeroNf           float32 `json:"numeroNf"`
	ValorTotal         float32 `json:"valorTotal"`
	DataNf             string  `json:"dataNf"`
	CnpjEmissorNf      string   `json:"cnpjEmissorNf"`
	CnpjDestinatarioNf string   `json:"cnpjDestinatarioNf"`
}

func NewNota(
	NumeroNf float32,
	ValorTotal float32,
	DataNf string,
	CnpjEmissorNf string,
	CnpjDestinatarioNf string,
) *Nota {
	return &Nota{
		Id:                 uuid.NewString(),
		NumeroNf:           NumeroNf,
		ValorTotal:         ValorTotal,
		DataNf:             DataNf,
		CnpjEmissorNf:      CnpjEmissorNf,
		CnpjDestinatarioNf: CnpjDestinatarioNf,
	}
}
