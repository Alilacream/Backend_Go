package main

import (
	"fmt"
	"log"
	"net/http"
)

type RequestData struct {
	Username string `json:"username"`
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world :(")
}

func main() {
	fmt.Println("Listening in Port 8080")
	http.HandleFunc("/", helloWorld)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
