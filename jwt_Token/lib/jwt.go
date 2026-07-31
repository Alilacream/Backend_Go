package lib

import (
	"alilacream/jwt/config"
	"alilacream/jwt/model"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateNewToken(user *model.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, model.User{
		RegisteredClaims: jwt.RegisteredClaims{},
		ID:               user.ID,
		Name:             user.Name,
		Password:         user.Password,
	})
	jwtToken, err := token.SignedString(config.GetEnv("SECRET"))
	if err != nil {
		return "", err
	}
	return jwtToken, nil
}
