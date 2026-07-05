package main

import (
	"net/http"
	"simple_web_server/routes"
)

func main() {
	root := http.Dir("./Simple_Web_Server")
	fileServer := http.FileServer(root)
	routes.Routing(&fileServer)
}
