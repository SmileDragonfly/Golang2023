package main

import (
	"bytes"
	"log"
)

func main() {
	blockSize := 8
	src := []byte{1, 2, 3}
	log.Println(src)
	padding := blockSize - len(src)%blockSize
	log.Println(byte(padding))
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	src = append(src, padtext...)
	log.Println(src)
}
