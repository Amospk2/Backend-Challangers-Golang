package product

type ProductRepository interface {
	Get() ([]Product, error)
	GetById(id string) (Product, error)
	Update(data Product) error
	Create(data Product) error
	Delete(id string) error
}
