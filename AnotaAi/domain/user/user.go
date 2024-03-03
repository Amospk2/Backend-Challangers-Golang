package user

type Users struct {
	Id       string `json:"id,omitempty"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Age      string `json:"age" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func NewUser(
	id string,
	name string,
	email string,
	age string,
	password string,
) *Users {
	return &Users{
		Id:       id,
		Name:     name,
		Email:    email,
		Password: password,
		Age:      age,
	}

}
