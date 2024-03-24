package nota

import (
	"fmt"
	"time"
)

func (n *Nota) Valid() bool {

	if fmt.Sprintf("%T", n.NumeroNf) != "float32" || n.NumeroNf < 0 {
		return false
	}


	if _, err := time.Parse("2006-01-02", n.DataNf); err != nil {
		return false
	}

	if fmt.Sprintf("%T", n.ValorTotal) != "float32" || n.ValorTotal < 0 {
		return false
	}

	if len(n.CnpjEmissorNf) == 0 || n.CnpjEmissorNf == "" {
		return false
	}

	if len(n.CnpjDestinatarioNf) == 0 || n.CnpjDestinatarioNf == "" {
		return false
	}

	return true
}
