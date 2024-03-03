package product

type Product struct {
	Id          string  `json:"id,omitempty"`
	Title       string  `json:"title,omitempty"`
	Description string  `json:"description" validate:"required"`
	Price       float32 `json:"price" validate:"required"`
	Category    string  `json:"category" validate:"required"`
	OwnerID     string  `json:"ownerID" validate:"required"`
}

func NewProduct(
	title string,
	description string,
	price float32,
	category string,
	ownerID string,
) *Product {
	return &Product{
		Title:       title,
		Description: description,
		Price:       price,
		Category:    category,
		OwnerID:     ownerID,
	}

}
