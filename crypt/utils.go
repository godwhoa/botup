package crypt

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/minio/blake2b-simd"
	"golang.org/x/crypto/bcrypt"
)

func RandomBytes() []byte {
	b := make([]byte, 20)
	rand.Read(b)
	return b
}

func UID(user string) (string, error) {
	hash, err := blake2b.New(&blake2b.Config{Size: 16})
	if err != nil {
		return "", err
	}
	b := []byte(user)
	hash.Write(b)
	hash.Write(RandomBytes())
	return hex.EncodeToString(hash.Sum(nil)), nil
}

func Verify(hashed string, pass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(pass))
	return err == nil
}

func Hash(pass string) (string, error) {
	// DefaultCost = 10
	hashed, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}
