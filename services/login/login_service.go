package login

import (
	"dashcode/general"
	"dashcode/repositories/loginrepo"
	"dashcode/services"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func Login(email, pwd string) (int64, *services.Error) {

	id, hash, err := loginrepo.FetchLogin(email)

	if err != nil {
		general.PotencialInternalError(err)
		return -1, services.ErrorInternal
	}

	if id == -1 {
		log.Println("Hey whatsup?")
		return -1, services.ErrorNotFound
	}

	err = bcrypt.CompareHashAndPassword(hash, []byte(pwd))

	if err != nil {
		return -1, &services.Error{
			Code: http.StatusNotFound,
			Err:  "password incorrect",
		}
	}

	return id, nil
}
