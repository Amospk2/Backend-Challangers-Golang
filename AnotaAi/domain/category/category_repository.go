package category

type CategoryRepository interface {
	Get() ([]Category, error)
	GetById(id string) (Category, error)
	Update(data Category) error
	Create(data Category) error
	Delete(id string) error
}
