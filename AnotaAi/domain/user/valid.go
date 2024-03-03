package user

import "time"

func (user Users) Valid() bool {
	if len(user.Name) == 0 && user.Name == "" {
		return false
	}

	if len(user.Email) == 0 && user.Email == "" {
		return false
	}

	if _, err := time.Parse("2006-01-02", user.Age); err != nil {
		return false
	}

	return true
}
