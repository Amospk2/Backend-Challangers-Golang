package user

func (user Users) Valid() bool {
	if len(user.Name) == 0 && user.Name == "" {
		return false
	}

	if len(user.Email) == 0 && user.Email == "" {
		return false
	}

	return true
}
