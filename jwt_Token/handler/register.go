package handler

import (
	"alilacream/jwt/model"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func Register(user *map[string]model.User) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Couldn't Parse the Form", http.StatusBadRequest)
			return
		}
		username := r.FormValue("name")
		password := r.FormValue("password")

		if len(username) < 8 || len(password) < 8 {
			http.Error(w, "Username length or Password length need to be at least 8 chars long", http.StatusConflict)
			return
		}
		// creates a hash
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 6)
		if err != nil {
			http.Error(w, "Unproccesed Password", http.StatusInternalServerError)
			return
		}
		NewUser := model.User{
			ID:       uuid.New().String(),
			Name:     username,
			Password: string(hashedPassword),
		}
		(*user)[username] = NewUser
		json.NewEncoder(w).Encode(NewUser)
	}
}
