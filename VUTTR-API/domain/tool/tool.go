package tool

type Tool struct {
	Id          string   `json:"id,omitempty"`
	Title       string   `json:"title" validate:"required,max=32"`
	Description string   `json:"description" validate:"required"`
	Link        string   `json:"link" validate:"required"`
	Tags        []string `json:"tags" validate:"dive,max=32"`
}

func NewTool(
	Id string,
	Title string,
	Description string,
	Link string,
	Tags []string,
) *Tool {
	return &Tool{
		Id:          Id,
		Title:       Title,
		Description: Description,
		Link:        Link,
		Tags:        Tags,
	}

}
