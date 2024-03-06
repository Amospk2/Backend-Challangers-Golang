package user

type UserRepository interface {
	Get() ([]Users, error)
	GetById(id string) (Users, error)
	GetByEmail(email string) (Users, error)
	Update(data Users) error
	Create(data Users) error
	Delete(id string) error
}
