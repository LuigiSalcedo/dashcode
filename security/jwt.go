package security

import (
	"dashcode/general"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

// Secret - The secret is on plain text just for academic purpose
var jwtSecret = "dashcode-secret-jwt"

// Possibles errors
var (
	ErrMapClaims = errors.New("error parsing from jwt to map claims")
)

type JWToken struct {
	Token string `json:"jwt"`
}

type TokenData struct {
	jwt.StandardClaims
	Id int64 `json:"id"`
}

// Generate a JWT using the USER ID
func GenerateJWT(id int64) *JWToken {

	data := TokenData{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(2 * time.Hour).Unix(),
		},
		id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, data)

	signedToken, err := token.SignedString([]byte(jwtSecret))

	if err != nil {
		general.PotencialInternalError(err)
		return nil
	}

	return &JWToken{Token: signedToken}
}

// Get the data from a JWT
func IDFromJWT(token string) (int64, error) {
	data, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return -1, err
	}

	dataMap, ok := data.Claims.(jwt.MapClaims)

	if !ok {
		general.PotencialInternalError(ErrMapClaims)
		return -1, nil
	}

	return int64(dataMap["id"].(float64)), nil
}
