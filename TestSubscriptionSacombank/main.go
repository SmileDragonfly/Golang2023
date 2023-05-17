package main

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/google/go-querystring/query"
	"github.com/google/uuid"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"
)

type SubData struct {
	ProfileID           string `json:"ProfileID"`
	AccessKey           string `json:"AccessKey"`
	TransactionID       string `json:"TransactionID"`
	TransactionDateTime string `json:"TransactionDateTime"`
	Language            string `json:"Language"`
	IsTokenRequest      string `json:"IsTokenRequest"`
	SubscribeOnly       string `json:"SubscribeOnly"`
	SubscribeWithMin    string `json:"SubscribeWithMin"`
	CardNumber          string `json:"CardNumber"`
	Description         string `json:"Description"`
	TotalAmount         string `json:"TotalAmount"`
	InternationalFee    string `json:"InternationalFee"`
	DomesticFee         string `json:"DomesticFee"`
	Currency            string `json:"Currency"`
	SSN                 string `json:"SSN"`
	FirstName           string `json:"FirstName"`
	LastName            string `json:"LastName"`
	Gender              string `json:"Gender"`
	Address             string `json:"Address"`
	District            string `json:"District"`
	City                string `json:"City"`
	PostalCode          string `json:"PostalCode"`
	Country             string `json:"Country"`
	Mobile              string `json:"Mobile"`
	Email               string `json:"Email"`
	CancelUrl           string `json:"CancelUrl"`
	ReturnUrl           string `json:"ReturnUrl"`
	Signature           string `json:"Signature"`
}

func main() {
	// Marshal data
	byte, err := os.ReadFile("Data.json")
	if err != nil {
		log.Fatalln(err)
	}
	var subData SubData
	err = json.Unmarshal(byte, &subData)
	if err != nil {
		log.Fatalln(err)
	}
	subData.TransactionID = GetTransactionID()
	subData.TransactionDateTime = GetTransactionDateTime()
	// Marshal data
	byteSubData, err := json.Marshal(subData)
	if err != nil {
		log.Fatalln(err)
	}
	// Unmarshal to map
	var subDataMap map[string]string
	err = json.Unmarshal(byteSubData, &subDataMap)
	// Get data to sign
	//sSign := GetStringFromMap(subDataMap)
	sSign := GetStringSign(subData)
	sSignature, err := CreateSignatureMD5("MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCVEi/qLwGQWsJk+BH9vrcavcrwaV/m8n4adWt0MPhlWvEqv7bcLoWYu2o+3whAiLY4ZybN6acfCqttAs2+rHcF88unh6w1PLqaCB5JErdeRmgXnYlCu6jIZMYhhvpl5yBfCTHDB1FGF47Bd/EfztIu21NScXddsN0tiaB1uT3/RPon68wokvEzOMMQl89Ik79XDIJPhBnsNU9GX838NaD3Chzdn7eFGUqUdOjbyoIVkHrXj4qfx1935YLX7d/g9auFr7k3ClE7iXJjRCJHmkF8c0seyD2N4TzYhPyAj7Ku76Uls83Tu7KrmQEkPUL2tRGyRC6wimnOlBPA1Ad0BCtPAgMBAAECggEADh6gYjacl8847XZfweCQGFzUvYvFlSrvzdLEYEeJJ5SDFlD9YYKYjdxmlljqv64TUGlY0BUrCrIohZVH5qEQTwSGDDK6GXrMk+1j7Pj0XF4f2ujFiQgcVwrQh3lw+zj2pHnK+FWXmaN+lo2lTaV53A1TKZJsw3eOo5MPG0d1+1kotGrZeOFELRnjLkGFw4SntahhMHByXKAl8pFGUyRJqHZBXw4erXbRLMDg7TAYn2VJdHrhSeEzFGmKA7QHuTRq52SZg71xDLiHmFVSeP7tVrj60jN7pjJTjJR6C2fFxphEpKaQ02zRzTjmqsAMbP77GmlWtyP/+crw70P4q/yHXQKBgQDFO0xn32CwlsTb+QIZnhfYO46+lfJP6MFuK/WIu0p2tee8xGDQYkWZ4xETkJb604ArhQ9WATDKfZkloCtk8pcRqVp6Q0xFjL+pfnZ4W9CUOwxpf1kPn8u6Dcay5CYLtbibrbaA5XR2HjsdxE+lxl29nfwabYEqCCmmPEPF2YaPZQKBgQDBfTrfa59KLzRHGAz55YXk6KNcjHzzRYT4b7Y4Hzh1OCB27tnVZjGBp8IUrXSnodPmFhcia/dQJJCZdciQ8Z8oj3XuyY4iB0QsXxBmFomjH/FYI+1hnhQehQvAqAeIpYlNRXYYPxqAfn3AQJOZebDrT/V4JaS4g9YjcQT4IqiGowKBgQCqZKeG2cIj7a1XSZJZ5W4+Pn39A3hbNv/dmZa/sOcNFeyF9baacTwmTbiUCYeWXSDO+F6ec9reJZIoom67AKYo+QGUvQ1ozMdMvFfHdbMGTNlVT1L3H5uXOo2eQWLpHO7HeFVCmHl8DnQOLGqPEogr6BBEGLTNRk4NMuVuSZZpzQKBgET23MgLdRAc+RYp9V4QuAOaA7gV/uc6rSVbs+gXAKmPIsshYRUVwqmC4MM7++tP29YTo5VKRDEVh1CbUayP4nmzgIZm4rkwO9VQ4OhyOgaheQVAcPitPmCObVzyxxSmY+Td0DTeMRUBgNLIcZNvc2a77jMvv6FgpC+ntey3dbffAoGAPvAnfdpFlUlRDdbsgn6HeiXcErCETeUPnCuiMjrb0zI1mnaJEPTJukQEFoHNmtFNgK4t1l0UEApYJ1PPPcyKVmvJSVwaXVxnLHvIj0vpNjTs9gdG37EZHFKXQjjhn+TD2iOXv1kqnyLJQjFBw+xgDtzp8qITXmJt3qKdpA7KQBw=", sSign)
	if err != nil {
		log.Fatalln(err)
	}
	err = VerifySignatureMD5RSA(sSignature, sSign, "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAlRIv6i8BkFrCZPgR/b63Gr3K8Glf5vJ+GnVrdDD4ZVrxKr+23C6FmLtqPt8IQIi2OGcmzemnHwqrbQLNvqx3BfPLp4esNTy6mggeSRK3XkZoF52JQruoyGTGIYb6ZecgXwkxwwdRRheOwXfxH87SLttTUnF3XbDdLYmgdbk9/0T6J+vMKJLxMzjDEJfPSJO/VwyCT4QZ7DVPRl/N/DWg9woc3Z+3hRlKlHTo28qCFZB614+Kn8dfd+WC1+3f4PWrha+5NwpRO4lyY0QiR5pBfHNLHsg9jeE82IT8gI+yru+lJbPN07uyq5kBJD1C9rURskQusIppzpQTwNQHdAQrTwIDAQAB")
	if err != nil {
		log.Fatalln(err)
	}
	subData.Signature = sSignature
	PrintIndent(subData)
	link := "https://cardtest.Sacombank.com.vn:9448/checkout/stbCheckout"
	//link := "https://87f6-125-235-61-74.ngrok-free.app/return"
	//link := "https://cardtest.sacombank.com.vn:9448/vtis/v1/retrieveStepUpMethods"
	form, err := query.Values(subData)
	log.Printf("Request: %+v", subData)
	//resp, err := http.PostForm(link, form)
	hc := http.Client{CheckRedirect: func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}}
	req, err := http.NewRequest("POST", link, strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := hc.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	// Parse response data
	respBodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	urlRet := "https://cardtest.Sacombank.com.vn:9448/checkout/" + resp.Header.Get("Location")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Subscription Response:", string(respBodyBytes))
	log.Println("Subscription Url:", urlRet)
	log.Println("Stop main")
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

func VerifySignatureMD5RSA(sSignature string, sRawData string, sPublicKey string) error {
	// Byte signature
	byteSignature, err := base64.StdEncoding.DecodeString(sSignature)
	if err != nil {
		return err
	}
	// Decode base64 signature
	bytePublicKey, err := base64.StdEncoding.DecodeString(sPublicKey)
	if err != nil {
		return err
	}
	publicKey, err := x509.ParsePKIXPublicKey(bytePublicKey)
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
func GetTransactionDateTime() string {
	t := time.Now()
	return t.UTC().Format("2006-01-02T03:04:05Z")
}

func GetTransactionID() string {
	id := uuid.New().String()
	return string([]byte(id)[:32])
}

func GetStringFromMap(mapData map[string]string) string {
	var sRet string
	mk := make([]string, len(mapData))
	i := 0
	for k, _ := range mapData {
		mk[i] = k
		i++
	}
	sort.Strings(mk)
	for _, value := range mk {
		sRet += mapData[value]
	}
	return sRet
}

func GetStringSign(des SubData) string {
	var sSigned string
	sSigned += des.AccessKey
	sSigned += des.Address
	sSigned += des.CancelUrl
	sSigned += des.City
	sSigned += des.Country
	sSigned += des.Currency
	sSigned += des.Description
	sSigned += des.District
	sSigned += des.DomesticFee
	sSigned += des.Email
	sSigned += des.FirstName
	sSigned += des.Gender
	sSigned += des.InternationalFee
	sSigned += des.IsTokenRequest
	sSigned += des.Language
	sSigned += des.LastName
	sSigned += des.Mobile
	sSigned += des.PostalCode
	sSigned += des.ProfileID
	sSigned += des.ReturnUrl
	sSigned += des.SSN
	sSigned += des.SubscribeOnly
	sSigned += des.SubscribeWithMin
	sSigned += des.TotalAmount
	sSigned += des.TransactionDateTime
	sSigned += des.TransactionID
	return sSigned
}

func PrintIndent(data any) {
	byteData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return
	}
	log.Println(string(byteData))
}
