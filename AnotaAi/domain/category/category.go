package category

type Category struct {
	Id          string `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"name" validate:"required"`
	OwnerID     string `json:"ownerID" validate:"required"`
}

func NewCategory(
	title string,
	description string,
	ownerID string,
) *Category {
	return &Category{
		Title:       title,
		Description: description,
		OwnerID:     ownerID,
	}

}
