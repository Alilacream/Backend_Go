package models

// only login user
type Login struct {
	HashedPassword string
	SessionToken   string
	CSRFToken      string
}
