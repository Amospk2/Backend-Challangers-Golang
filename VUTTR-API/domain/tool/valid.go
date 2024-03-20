package tool

func (tool Tool) Valid() bool {

	if len(tool.Title) == 0 && tool.Title == "" {
		return false
	}

	if len(tool.Description) == 0 && tool.Description == "" {
		return false
	}

	if len(tool.Tags) == 0 {
		return false
	}

	return true
}
