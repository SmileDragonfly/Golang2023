package Utils

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
)

func PKCS5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}

func EncryptTripleDESB64(sSrc string, sKey string) (string, error) {
	byteSrc := []byte(sSrc)
	byteKey := []byte(sKey)
	block, err := des.NewTripleDESCipher(byteKey)
	if err != nil {
		return "", err
	}
	iv := byteKey[:8]
	mode := cipher.NewCBCEncrypter(block, iv)
	byteSrc = PKCS5Padding(byteSrc, block.BlockSize())
	byteDes := make([]byte, len(byteSrc))
	mode.CryptBlocks(byteDes, byteSrc)
	sDes := base64.StdEncoding.EncodeToString(byteDes)
	return sDes, nil
}

func DecryptTripleDESB64(sSrc string, sKey string) (string, error) {
	byteSrc, err := base64.StdEncoding.DecodeString(sSrc)
	if err != nil {
		return "", err
	}
	byteKey := []byte(sKey)
	block, err := des.NewTripleDESCipher(byteKey)
	if err != nil {
		return "", err
	}
	iv := byteKey[:8]
	mode := cipher.NewCBCDecrypter(block, iv)
	byteDes := make([]byte, len(byteSrc))
	mode.CryptBlocks(byteDes, byteSrc)
	byteDes = PKCS5UnPadding(byteDes)
	sDes := string(byteDes)
	return sDes, nil
}
