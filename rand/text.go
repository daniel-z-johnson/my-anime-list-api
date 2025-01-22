package rand

import (
	"crypto/rand"
	"encoding/base64"
)

func State() (string, error) {
	bites := make([]byte, 36)
	_, err := rand.Read(bites)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bites), nil
}
