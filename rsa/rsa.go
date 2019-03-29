package rsa

import (
	"crypto/rand"
	"crypto/rsa"
)

func genkey(publicKeyPath, privateKeyPath string, size int) {
	privateKey, err := rsa.GenerateKey(rand.Reader, size)
	if err != nil {

	}
}
