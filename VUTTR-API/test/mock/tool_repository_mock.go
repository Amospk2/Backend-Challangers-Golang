package mock

import (
	"VUTTR-API/domain/tool"
	"errors"
	"strings"
)

type ToolRepositoryMock struct {
	datas []tool.Tool
}

func (db *ToolRepositoryMock) Get() ([]tool.Tool, error) {
	return db.datas, nil
}

func (db *ToolRepositoryMock) GetByTag(tag string) ([]tool.Tool, error) {
	tools := make([]tool.Tool, 0)

	for _, content := range db.datas {

		for _, c := range content.Tags {
			if strings.Compare(c, tag) == 0 {
				tools = append(tools, content)
				break
			}
		}
	}

	return tools, nil
}

func (db *ToolRepositoryMock) GetById(id string) (tool.Tool, error) {
	var idx int = -1

	for index, content := range db.datas {
		if id == content.Id {
			idx = index
		}
	}

	if idx < 0 {
		return tool.Tool{}, errors.New("NOT FOUND")
	}

	return db.datas[idx], nil
}

func (db *ToolRepositoryMock) Update(data tool.Tool) error {
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

func (db *ToolRepositoryMock) Create(data tool.Tool) error {
	db.datas = append(db.datas, data)

	return nil
}

func (db *ToolRepositoryMock) Delete(id string) error {
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

func NewToolRepositoryMock(users []tool.Tool) *ToolRepositoryMock {
	return &ToolRepositoryMock{datas: users}
}
