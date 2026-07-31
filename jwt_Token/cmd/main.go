package main

import (
	"log"
	"net/http"

	"alilacream/jwt/cmd/api"
)

func main() {
	mux := http.NewServeMux()
	app := api.NewServerApi(":8080", mux)
	log.Println("Server Listening in Port 8080...")
	// routes
	app.Routes()
	log.Fatal(app.Run())
}
