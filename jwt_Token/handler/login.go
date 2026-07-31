package handler

import (
	"alilacream/jwt/lib"
	"alilacream/jwt/model"
	"encoding/json"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func Login(user *map[string]model.User) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer w.Write([]byte("Logged In!"))

		var formValues model.User
		w.Header().Set("Content/Type", "application/json")

		json.NewDecoder(r.Body).Decode(&formValues)
		if formValues.Name == "" || formValues.Password == "" {
			http.Error(w, "You have to set username/password", http.StatusMethodNotAllowed)
			return
		}
		if _, ok := (*user)[formValues.Name]; !ok {
			http.Error(w, "User Does not Exist!", http.StatusBadRequest)
			return
		}
		if err := bcrypt.CompareHashAndPassword([]byte((*user)[formValues.Name].Password), []byte(formValues.Password)); err != nil {
			http.Error(w, "Password False", http.StatusConflict)
			return
		}
		// setting up a jwt token
		jwtToken, err := lib.GenerateNewToken(&formValues)
		if err != nil {
			http.Error(w, "Couldn't generate Token", http.StatusInternalServerError)
			return
		}
		// it's propably vuln to csrf attack
		// since jwt tokens are stored in http Cookie, i can use it for localStorage with Secure attribute set as true and same origin policy
		// => result no vulnerability
		http.SetCookie(w, &http.Cookie{
			Name:     "Jwt-Token",
			Value:    jwtToken,
			Expires:  time.Now().Add(24 * time.Hour),
			HttpOnly: true,
		})
	}
}
