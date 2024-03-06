package mock

import (
	"api/domain/product"
	"errors"
)

type ProductRepositoryMock struct {
	datas []product.Product
}

func (db *ProductRepositoryMock) Get() ([]product.Product, error) {
	return db.datas, nil
}

func (db *ProductRepositoryMock) GetById(id string) (product.Product, error) {
	var idx int = -1

	for index, content := range db.datas {
		if id == content.Id {
			idx = index
		}
	}

	if idx < 0 {
		return product.Product{}, errors.New("NOT FOUND")
	}

	return db.datas[idx], nil
}

func (db *ProductRepositoryMock) GetByEmail(email string) (product.Product, error) {

	return product.Product{}, nil
}

func (db *ProductRepositoryMock) Update(data product.Product) error {
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

func (db *ProductRepositoryMock) Create(data product.Product) error {
	db.datas = append(db.datas, data)

	return nil
}

func (db *ProductRepositoryMock) Delete(id string) error {
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

func NewProductRepositoryMock(products []product.Product) *ProductRepositoryMock {
	return &ProductRepositoryMock{datas: products}
}
