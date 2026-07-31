package handler

import (
	"alilacream/jwt/model"
	"net/http"
)

func logout(user *map[string]model.User) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		//?
	}
}
