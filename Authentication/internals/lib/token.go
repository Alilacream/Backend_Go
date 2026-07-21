package lib

import (
	"crypto/rand"
	"encoding/base64"
	"log"
)

func GenerateTok(length int) string {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		log.Fatalln("[ERROR]: failed to generate a new token: ", err)
	}
	return base64.URLEncoding.EncodeToString(bytes)
}
