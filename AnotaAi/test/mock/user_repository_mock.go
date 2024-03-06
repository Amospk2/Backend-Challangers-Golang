package mock

import (
	"api/domain/user"
	"errors"
)

type UserRepositoryMock struct {
	datas []user.Users
}

func (db *UserRepositoryMock) Get() ([]user.Users, error) {
	return db.datas, nil
}

func (db *UserRepositoryMock) GetById(id string) (user.Users, error) {
	var idx int = -1

	for index, content := range db.datas {
		if id == content.Id {
			idx = index
		}
	}

	if idx < 0 {
		return user.Users{}, errors.New("NOT FOUND")
	}

	return db.datas[idx], nil
}

func (db *UserRepositoryMock) GetByEmail(email string) (user.Users, error) {
	var idx int = -1

	for index, content := range db.datas {
		if email == content.Email {
			idx = index
		}
	}

	if idx < 0 {
		return user.Users{}, errors.New("NOT FOUND")
	}

	return db.datas[idx], nil
}

func (db *UserRepositoryMock) Update(data user.Users) error {
	var idx int = -1

	for index, content := range db.datas {
		if data.Id == content.Id {
			idx = index
		}
	}

	if idx < 0 {
		return errors.New("NOT FOUND")
	}

	db.datas[idx] = data

	return nil
}

func (db *UserRepositoryMock) Create(data user.Users) error {
	db.datas = append(db.datas, data)

	return nil
}

func (db *UserRepositoryMock) Delete(id string) error {
	var idx int = -1

	for index, content := range db.datas {
		if id == content.Id {
			idx = index
		}
	}

	if idx < 0 {
		return errors.New("NOT FOUND")
	}

	db.datas = append(db.datas[:idx], db.datas[idx+1:]...)

	return nil
}

func NewUserRepositoryMock(users []user.Users) *UserRepositoryMock {
	return &UserRepositoryMock{datas: users}
}
