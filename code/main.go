package main

import (
	cryptoRand "crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"
)

func generateCodeVerifier() string {
	randomBytes := make([]byte, 32)
	_, err := cryptoRand.Read(randomBytes)
	if err != nil {
		panic(err)
	}
	return strings.TrimRight(base64.URLEncoding.EncodeToString(randomBytes), "=")
}

func generateCodeChallenge(codeVerifier string) string {
	hashed := sha256.Sum256([]byte(codeVerifier))
	return strings.TrimRight(base64.URLEncoding.EncodeToString(hashed[:]), "=")
}

func main() {
	codeVerifier := generateCodeVerifier()
	codeChallenge := generateCodeChallenge(codeVerifier)

	fmt.Printf("Code Verifier: %s\n", codeVerifier)
	fmt.Printf("Code Challenge (S256): %s\n", codeChallenge)
}
