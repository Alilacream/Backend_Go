package main

import (
	"fmt"
	"net/http"
)

func routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my server")
	})
	return rateLimiter(mux, 2, 10)
}

func main() {
	srv := routes()
	http.ListenAndServe(":8080", srv)
}
