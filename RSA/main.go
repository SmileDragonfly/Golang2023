package main

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/pem"
	"log"
	"os"
)

func main() {
	pubKeyByte, err := os.ReadFile("./PublicKey.pem")
	if err != nil {
		log.Println(err.Error())
	}
	//log.Println(pubKeyByte)
	pubBlock, _ := pem.Decode(pubKeyByte)
	if pubBlock == nil || pubBlock.Type != "PUBLIC KEY" {
		log.Fatal("Failed to decode PEM block containing public key")
	}
	pub, err := x509.ParsePKIXPublicKey(pubBlock.Bytes)
	if err != nil {
		log.Fatal("Parse public key failed", err.Error())
	}
	log.Println("Public key:", pub)
	rsaPubKey, _ := pub.(*rsa.PublicKey)
	// PRIVATE KEY
	priKeyByte, err := os.ReadFile("./PrivateKey.pem")
	if err != nil {
		log.Println(err.Error())
	}
	//log.Println(priKeyByte)
	priBlock, _ := pem.Decode(priKeyByte)
	if priBlock == nil || priBlock.Type != "RSA PRIVATE KEY" {
		log.Fatal("Failed to decode PEM block containing private key")
	}
	pri, err := x509.ParsePKCS1PrivateKey(priBlock.Bytes)
	if err != nil {
		log.Fatal("Parse private key failed", err.Error())
	}
	log.Println("Private key:", pri)
	//publicKey, err := base64.StdEncoding.DecodeString(publicKeyB64)
	//if err != nil {
	//	log.Println(err.Error())
	//}
	//
	//log.Println("Public key: ", publicKey)
	//privateKey, err := base64.StdEncoding.DecodeString(privateKeyB64)
	//if err != nil {
	//	log.Println(err.Error())
	//}
	//log.Println("Private key: ", privateKey)
	data := []byte("Toi la doan trong dat")
	md5Data := md5.Sum(data)
	sha1Data := sha1.Sum(md5Data[:])
	signatureData, err := rsa.SignPKCS1v15(rand.Reader, pri, crypto.SHA1, sha1Data[:])
	if err != nil {
		log.Fatal("Signature data failed: ", err.Error())
	}
	log.Println("Signature data:", signatureData)
	err = rsa.VerifyPKCS1v15(rsaPubKey, crypto.SHA1, sha1Data[:], signatureData)
	if err != nil {
		log.Fatal("Verify signature failed: ", err.Error())
	}
	log.Println("Verify signature success")
}
