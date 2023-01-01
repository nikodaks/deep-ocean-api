package hash

import (
	"crypto/sha1"
	"fmt"
)

type SHA1Hasher struct {
	salt string
}

func NewSHA1Hasher(salt string) *SHA1Hasher {
	return &SHA1Hasher{salt: salt}
}

func (h *SHA1Hasher) Hash(data string) (string, error) {
	hasher := sha1.New()

	if _, err := hasher.Write([]byte(data)); err != nil {
		return "", err
	}

	hash := fmt.Sprintf("%x", hasher.Sum([]byte(h.salt)))
	return hash, nil
}
