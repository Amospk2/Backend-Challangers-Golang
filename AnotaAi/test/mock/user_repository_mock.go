package mock

import "api/domain/user"

type UserRepositoryMock struct {
	datas []user.Users
}

func (db UserRepositoryMock) Get() ([]user.Users, error) {
	return db.datas, nil
}

func (db UserRepositoryMock) GetById(id string) (user.Users, error) {

	return user.Users{}, nil
}

func (db UserRepositoryMock) GetByEmail(email string) (user.Users, error) {

	return user.Users{}, nil
}

func (db UserRepositoryMock) Update(data user.Users) error {

	return nil
}

func (db UserRepositoryMock) Create(data user.Users) error {

	return nil
}

func (db UserRepositoryMock) Delete(id string) error {

	return nil
}


func NewUserRepositoryMock(users []user.Users) *UserRepositoryMock {
	return &UserRepositoryMock{datas: users}
}