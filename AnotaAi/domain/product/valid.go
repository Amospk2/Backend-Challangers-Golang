package product

import "fmt"

func (product Product) Valid() bool {
	if len(product.Title) == 0 && product.Title == "" {
		return false
	}

	if len(product.Description) == 0 && product.Description == "" {
		return false
	}

	if fmt.Sprintf("%T", product.Price) != "int" && product.Price < 0 {
		return false
	}
	
	return true
}
