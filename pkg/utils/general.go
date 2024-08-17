package utils

import (
	"crypto/rand"
	"encoding/base64"

	"github.com/google/uuid"
)

func GenerateUuid() string {
	return uuid.New().String()
}

func GenerateRandomString(length int) (string, error) {
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	randomString := base64.URLEncoding.EncodeToString(randomBytes)

	return randomString[:length], nil
}
