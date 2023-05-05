package main

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"errors"
	"time"
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

func ReadPublicKey(sPublicKey string) (any, error) {
	bytePublicKey, err := base64.StdEncoding.DecodeString(sPublicKey)
	if err != nil {
		return nil, err
	}
	publicKey, err := x509.ParsePKIXPublicKey(bytePublicKey)
	if err != nil {
		return nil, err
	}
	return publicKey, nil
}

func CreateSignatureMD5(sPrivateKey string, sRawData string) (string, error) {
	privateKey, err := ReadPrivateKey(sPrivateKey)
	if err != nil {
		return "", err
	}
	// MD5 data
	md5 := md5.Sum([]byte(sRawData))
	// Sign md5 data
	rsaPrivateKey, ok := privateKey.(*rsa.PrivateKey)
	if !ok {
		return "", errors.New("Invalid private key")
	}
	byteSignature, err := rsa.SignPKCS1v15(rand.Reader, rsaPrivateKey, crypto.MD5, md5[:])
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(byteSignature), nil
}

func VerifySignatureMD5(sSignature string, sRawData string, sPublicKey string) error {
	// Byte signature
	byteSignature, err := base64.StdEncoding.DecodeString(sSignature)
	if err != nil {
		return err
	}
	// Decode base64 signature
	publicKey, err := ReadPublicKey(sPublicKey)
	if err != nil {
		return err
	}
	// MD5 body data
	byteMD5 := md5.Sum([]byte(sRawData))
	rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		return errors.New("Invalid public key")
	}
	return rsa.VerifyPKCS1v15(rsaPublicKey, crypto.MD5, byteMD5[:], byteSignature)
}

// Get RequestDateTime: yyyy-MM-ddTHH:mm:ssZ (UTC DateTime)
func GetRequestDateTime() string {
	t := time.Now()
	return t.UTC().Format("2006-01-02T03:04:05Z")
}

func GenerateTransactionID() string {
	return ""
}
