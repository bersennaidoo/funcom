package hash

import (
	"crypto/sha256"
	"encoding/base64"
	"math/rand"
)

const randomLength = 16

func GenerateSalt() string {
	var salt []byte
	var asciiPad int64
	asciiPad = 32

	for i := 0; i < randomLength; i++ {
		salt = append(salt, byte(rand.Int63n(94)+asciiPad))
	}

	return string(salt)
}

func GenerateHash(salt string, password string) string {
	var hash string
	fullString := salt + password
	sha := sha256.New()
	sha.Write([]byte(fullString))
	hash = base64.URLEncoding.EncodeToString(sha.Sum(nil))

	return hash
}
