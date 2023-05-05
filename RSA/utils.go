package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"os"
)

func ReadPrivateKey(sPrivateKey string) (any, error) {
	bytePrivateKey, err := base64.StdEncoding.DecodeString(sPrivateKey)
	if err != nil {
		return nil, err
	}
	privateKey, err := x509.ParsePKCS8PrivateKey(bytePrivateKey)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

func ReadPrivateKeyFromFile(sPath string) (*rsa.PrivateKey, error) {
	bytePrivateKey, err := os.ReadFile(sPath)
	if err != nil {
		return nil, err
	}
	// Load private key
	pemBlock, _ := pem.Decode(bytePrivateKey)
	if pemBlock == nil || pemBlock.Type != "RSA PRIVATE KEY" {
		return nil, errors.New("Decode pem file failed")
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(pemBlock.Bytes)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}
