package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/alilacream/auth/internals/lib"
	"github.com/alilacream/auth/models"
)

var users = map[string]models.Login{}

func register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST methods are allowed", http.StatusMethodNotAllowed)
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")
	if len(password) < 8 || len(username) < 8 {
		http.Error(w, "Invalid Password or Username", http.StatusNotAcceptable)
		return
	}
	// if user already exists
	if _, ok := users[username]; ok {
		http.Error(w, "User already exists", http.StatusConflict)
		return
	}
	hashedPass, _ := lib.HashPassword(password)
	users[username] = models.Login{
		HashedPassword: hashedPass,
	}
	fmt.Fprintln(w, "Done created the user hh")
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST methods are allowed", http.StatusMethodNotAllowed)
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Couldn't process Request", http.StatusNotAcceptable)
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")
	user, ok := users[username]
	// conditions: one for user's existance, the other is for the password
	if !ok {
		http.Error(w, "User Does not Exist ", http.StatusConflict)
		return
	}

	if !lib.CheckPassword(user.HashedPassword, password) {
		http.Error(w, "Password False", http.StatusConflict)
		return
	}

	sessionToken := lib.GenerateTok(10)
	csrfToken := lib.GenerateTok(32)
	// setting the session token in every request the user sends
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    sessionToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true, // inaccesable via client side (no js manipulation)
	})
	// setting csrf token preventing malicious csrf attacks
	http.SetCookie(w, &http.Cookie{
		Name:     "csrf_token",
		Value:    csrfToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: false, // accesable when client side is rendered

	})
	user.SessionToken = sessionToken
	user.CSRFToken = csrfToken

	users[username] = user

	fmt.Fprintf(w, "You're logged!")
}

func logout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST methods are allowed", http.StatusMethodNotAllowed)
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Couldn't process Request", http.StatusNotAcceptable)
		return
	}
	username := r.FormValue("username")
	user, ok := users[username]
	if !ok {
		http.Error(w, "You've miss inputed you're name!", http.StatusConflict)
		return
	}
	user = models.Login{
		HashedPassword: user.HashedPassword,
		SessionToken:   "",
		CSRFToken:      "",
	}
	users[username] = user
	fmt.Fprintf(w, "%s logged out!", username)
}

func Routes() {
	http.HandleFunc("/register", register)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	//	http.HandleFunc("/protected",handler.protected)
	http.ListenAndServe(":8080", nil)
}

func main() {
	defer fmt.Println("[CLOSING]: Port 8080 is closed, ")
	fmt.Println("[START]: Staring the Server in port 8080")
	Routes()
}
