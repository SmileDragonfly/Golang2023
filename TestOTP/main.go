package main

import (
	"encoding/base32"
	"fmt"
	"github.com/google/uuid"
	"github.com/pquerna/otp/totp"
	"time"
)

func main() {
	key := uuid.New().String()
	code, err := totp.GenerateCode(base32.StdEncoding.EncodeToString([]byte(key)), time.Now())
	if err != nil {
		panic(err)
	}
	fmt.Println("Code:", code)
	bOk := totp.Validate(code, base32.StdEncoding.EncodeToString([]byte(key)))
	if bOk != true {
		panic(bOk)
	}
}
