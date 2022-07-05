package service

import (
	"crypto/sha1"
	"fmt"
)

const salt = "hfdsfbwjhevfvfbdvbds"

func generatePasswordHash(password string) string {
	hash := sha1.New()

	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
