package utils

import (
    "crypto/sha256"
    "encoding/hex"
)

func Sha256Encrytion(text string) string {
	message := []byte(text)
	hash := sha256.New()
    hash.Write(message)
    bytes := hash.Sum(nil)
    hashCode := hex.EncodeToString(bytes)
    return hashCode
}