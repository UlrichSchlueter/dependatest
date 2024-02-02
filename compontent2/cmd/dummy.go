package main

import (
	"crypto/rand"
	"fmt"

	"golang.org/x/crypto/argon2"
)

func main() {
	password := []byte("my-password")
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		panic(err)
	}

	key := argon2.IDKey(password, salt, 1, 64*1024, 4, 32)

	fmt.Printf("Derived key: %x\n", key)
}
