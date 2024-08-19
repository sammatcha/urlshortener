package utils

import (
	"crypto/rand"
	"encoding/base64"
	"strings"
)
func GenerateShortURL()(string, error) {
	b:= make([]byte, 8) //generate 8 random bytes
	if _,err := rand.Read(b); err != nil{
		return "", err
	}
	shortURL := base64.URLEncoding.EncodeToString(b)
	return strings.TrimRight(shortURL, "="), nil //remove padding characters
}