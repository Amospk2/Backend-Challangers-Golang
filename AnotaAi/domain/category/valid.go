package category

func (category Category) Valid() bool {

	if len(category.Title) == 0 && category.Title == "" {
		return false
	}

	if len(category.Description) == 0 && category.Description == "" {
		return false
	}

	return true
}
