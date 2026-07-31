package handlers

import (
	"alilacream/jwt/model"
	"errors"
	"log"
	"net/http"
	"runtime/debug"
)

func RecoveryMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Caught panic: %v, Stack Trace: %s\n", err, string(debug.Stack()))
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
		}()
		// call the next handler in the chain
		next(w, r)
	}
}

var AuthorizedError = errors.New("Unauthorized")

func AuthMiddlWare(user *map[string]model.User, r *http.Request) error {
	return nil
}
