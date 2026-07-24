package handlers

import (
	"log"
	"net/http"
	"runtime/debug"

	"github.com/alilacream/auth/errors"
)

func RecoveryMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Caught panic: %v, Stack Trace: %s\n", err, string(debug.Stack()))
				errors.DisplayErr(w, "Internal")
			}
		}()
		// call the next handler in the chain
		next(w, r)
	}
}
