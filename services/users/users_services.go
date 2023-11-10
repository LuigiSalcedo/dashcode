package users

import (
	"dashcode/repositories/usersrepo"
	"dashcode/services"

	"golang.org/x/crypto/bcrypt"
)

// Service to save an user in the database
func SaveUser(id int64, name, email, password string) *services.Error {
	pwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return services.ErrorInternal
	}

	err = usersrepo.SaveUser(id, name, email, pwd)

	if err != nil {
		return services.ErrorBadRequest
	}

	return nil
}
