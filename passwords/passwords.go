package passwords

import (
	"crypto/rand"
	"encoding/base64"
)

const PASSWORD_LENGTH = 512

func GetPassword(hostname string) (string, error) {

	return "", nil
}

func GenerateRandomPassword() (string, error) {
	bytes := make([]byte, PASSWORD_LENGTH)

	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(bytes), nil
}
