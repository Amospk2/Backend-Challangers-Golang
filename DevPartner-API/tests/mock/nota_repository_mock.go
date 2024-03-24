package mock

import (
	"devpartner-api/domain/nota"
	"errors"
)

type NotaRepositoryMock struct {
	datas []nota.Nota
}

func (db *NotaRepositoryMock) Get() ([]nota.Nota, error) {
	return db.datas, nil
}

func (db *NotaRepositoryMock) GetById(id string) (nota.Nota, error) {
	var idx int = -1

	for index, content := range db.datas {
		if id == content.Id {
			idx = index
		}
	}

	if idx < 0 {
		return nota.Nota{}, errors.New("NOT FOUND")
	}

	return db.datas[idx], nil
}

func (db *NotaRepositoryMock) Update(data nota.Nota) error {
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

func (db *NotaRepositoryMock) Create(data nota.Nota) error {
	db.datas = append(db.datas, data)

	return nil
}

func (db *NotaRepositoryMock) Delete(id string) error {
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

func NewNotaRepositoryMock(users []nota.Nota) *NotaRepositoryMock {
	return &NotaRepositoryMock{datas: users}
}
