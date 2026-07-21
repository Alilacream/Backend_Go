package lib

import (
	"crypto/rand"
	"encoding/base64"
	"log"

	"golang.org/x/crypto/bcrypt" //nolint
)

// hashing the password (power to the 5)
func HashPassword(plainText string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(plainText), 5)
	return string(hashedBytes), err
}

// checks the equality between the password and the hashed version
func CheckPassword(password, hashedOne string) bool {
	return bcrypt.CompareHashAndPassword([]byte(password), []byte(hashedOne)) == nil
}

func GenerateTok(length int) string {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		log.Fatalln("[ERROR]: failed to generate a new token: ", err)
	}
	return base64.URLEncoding.EncodeToString(bytes)
}
