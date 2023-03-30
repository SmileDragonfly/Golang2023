package main

import (
	"DES/Utils"
	"log"
)

func main() {
	secretKey := "TestSecretKey32Character"
	sSource := "Toi la Doan Trong Dat"
	sEncrypted, err := Utils.EncryptTripleDESB64(sSource, secretKey)
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println("Encrypted:", sEncrypted)
	sDecrypted, err := Utils.DecryptTripleDESB64(sEncrypted, secretKey)
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println("Decrypted:", sDecrypted)
}
