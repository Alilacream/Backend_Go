package models

import "net/http"

// only login user
type Login struct {
	HashedPassword string
	SessionToken   string
	CSRFToken      string
}

// for errors
type HTTPError struct {
	Writer  http.ResponseWriter
	ErrStat string
	Status  int16
}
