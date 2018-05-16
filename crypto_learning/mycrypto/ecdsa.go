package mycrypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"errors"
	"fmt"
)

func GenerateEccKey() (*ecdsa.PrivateKey, error){
	curve := elliptic.P256()

	privKey, err := ecdsa.GenerateKey(curve, rand.Reader)

	if err != nil {
		errors.New("gen key error")
	}

	fmt.Println(privKey)

	return privKey, nil
}
