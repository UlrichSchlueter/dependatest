package main

import (
	"fmt"

	"golang.org/x/crypto/blake2b"
)

func main() {
	data := []byte("my-data")

	// Hash the data
	hash, err := blake2b.New256(nil)
	if err != nil {
		panic(err)
	}
	hash.Write(data)
	digest := hash.Sum(nil)

	fmt.Printf("Digest: %x\n", digest)
}
