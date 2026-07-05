package routes

import (
	"fmt"
	"log"
	"net/http"
)

// FormHandler the form handler
func FormHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/form" {
		http.Error(w, "Page Not Found", http.StatusNotFound)
	}
	if r.Method != "POST" {
		http.Error(w, "Method Not allowed", http.StatusMethodNotAllowed)
	}
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Couldn't parse the form: %v", err)
		return
	}
	Name := r.FormValue("name")
	Email := r.FormValue("email")
	fmt.Fprintf(w, "Name: %s\n", Name)
	fmt.Fprintf(w, "Email: %s\n", Email)
}

// HelloHandler the hello world returner
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	// well perhaps this is new to me
	if r.URL.Path != "/hello" {
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}
	// now i understand why we use this.
	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
	fmt.Fprintf(w, "Welcome to My Static Page")
}

func Routing(f *http.Handler) {
	defer fmt.Println("Server At port 8080 rn :P")
	http.Handle("/", *f)
	http.HandleFunc("/form", FormHandler)
	http.HandleFunc("/hello", HelloHandler)
	fmt.Println("Starting Server At Port 8080 in the LocalHost Ip")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("[ERROR]: Couldn't Serve the Port 8080 in the Server: %v", err)
		return
	}
}
