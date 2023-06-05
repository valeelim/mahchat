package utils

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateOpaqueToken(length int) (string, error) {
	tokenBytes := make([]byte, length)
	if _, err := rand.Read(tokenBytes); err != nil {
		return "", err
	}	
	token := base64.RawURLEncoding.EncodeToString(tokenBytes)
	return token, nil
}
