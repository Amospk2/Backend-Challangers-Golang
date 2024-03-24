package nota

type NotaRepository interface {
	Get() ([]Nota, error)
	GetById(id string) (Nota, error)
	Update(data Nota) error
	Create(data Nota) error
	Delete(id string) error
}