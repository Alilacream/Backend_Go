package handlers

import (
	"errors"
	"net/http"

	"github.com/alilacream/auth/models"
)

var AuthErr = errors.New("Unauthorized")

func Authorize(users *map[string]models.Login, r *http.Request) error {
	username := r.FormValue("username")
	user, ok := (*users)[username]

	if !ok {
		return AuthErr
	}
	// session token unmatched
	st, err := r.Cookie("session_token")
	if err != nil || st.Value == "" || st.Value != user.SessionToken {
		return AuthErr
	}
	// csrf unmatched
	csrfT := r.Header.Get("X-CSRF-Token")
	if csrfT == "" || csrfT != user.CSRFToken {
		return AuthErr
	}
	return nil
}
