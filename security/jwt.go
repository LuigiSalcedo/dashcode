package security

import (
	"dashcode/general"

	"github.com/golang-jwt/jwt"
)

// Secret - The secret is on plain text just for academic purpose
var jwtSecret = "dashcode-secret-jwt"

type JWToken struct {
	Token string `json:"jwt"`
}

func GenerateJWT(id int64) *JWToken {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": id,
	})

	signedToken, err := token.SignedString([]byte(jwtSecret))

	if err != nil {
		general.PotencialInternalError(err)
		return nil
	}

	return &JWToken{Token: signedToken}
}
