package tool

type ToolRepository interface {
	Get() ([]Tool, error)
	GetById(id string) (Tool, error)
	GetByTag(tag string) ([]Tool, error)
	Update(data Tool) error
	Create(data Tool) error
	Delete(id string) error
}
