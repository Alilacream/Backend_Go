package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
	"strconv"
)

type Person struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var (
	people = []Person{}
	static = 0
)

func getNewId() int {
	static++
	return static
}

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to our server"))
}

func getUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query().Get("id")
	iid, err := strconv.Atoi(id)

	if err != nil || id == "" {
		http.Error(w, "Set an ID", http.StatusBadRequest)
		return
	}

	for _, person := range people {
		if person.ID == iid {
			json.NewEncoder(w).Encode(person)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		} else {
			fmt.Fprintf(w, "User does not exist")
		}
	}
}

// getting all users
func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(people)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// inserting a new user
func InsertUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var person Person
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// check if the user is there.
	if slices.Contains(people, person) {
		http.Error(w, "User already Exists", http.StatusMethodNotAllowed)
		return
	}
	person.ID = getNewId()
	people = append(people, person)
	w.Write([]byte("New user registered"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", welcome)
	mux.HandleFunc("POST /users", getUsers)
	mux.HandleFunc("POST /users/new", InsertUser)
	mux.HandleFunc("GET /users/{id}", getUserByID)
	http.ListenAndServe(":8080", mux)
}
