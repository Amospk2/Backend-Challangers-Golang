package mock

import (
	"api/domain/category"
	"errors"
)

type CategoryRepositoryMock struct {
	datas []category.Category
}

func (db *CategoryRepositoryMock) Get() ([]category.Category, error) {
	return db.datas, nil
}

func (db *CategoryRepositoryMock) GetById(id string) (category.Category, error) {
	var idx int = -1

	for index, content := range db.datas {
		if id == content.Id {
			idx = index
		}
	}

	if idx < 0 {
		return category.Category{}, errors.New("NOT FOUND")
	}

	return db.datas[idx], nil
}

func (db *CategoryRepositoryMock) GetByEmail(email string) (category.Category, error) {

	return category.Category{}, nil
}

func (db *CategoryRepositoryMock) Update(data category.Category) error {
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

func (db *CategoryRepositoryMock) Create(data category.Category) error {
	db.datas = append(db.datas, data)

	return nil
}

func (db *CategoryRepositoryMock) Delete(id string) error {
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

func NewCategoryRepositoryMock(products []category.Category) *CategoryRepositoryMock {
	return &CategoryRepositoryMock{datas: products}
}
