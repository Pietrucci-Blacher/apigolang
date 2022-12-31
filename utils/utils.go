package utils

import (
	"crypto/sha512"
	"encoding/hex"
)

func Sha512(strToHash string) string {
	data := []byte(strToHash)
	hash := sha512.Sum512(data)
	return hex.EncodeToString(hash[:])
}
