package model

import "github.com/golang-jwt/jwt/v5"

type User struct {
	jwt.RegisteredClaims
	ID       string `json:"id"`
	Name     string `json:"name"`
	Password string `json:"-"`
}
