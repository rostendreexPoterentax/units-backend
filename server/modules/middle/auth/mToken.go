package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rotisserie/eris"
)

type tokenStruct struct {
	UserId int    `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

var jwtKey = []byte(os.Getenv("jwt_secret"))

func GenerateToken(userId int, email string) (string, error) {
	lifeTime := time.Now().UTC().Add(240 * time.Hour)

	newTokenData := &tokenStruct{
		UserId: userId,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(lifeTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, newTokenData)
	tokenStr, er := token.SignedString(jwtKey)
	if er != nil {
		return "", eris.Wrap(er, "failed to generate tokenStr")
	}
	return tokenStr, nil

}
