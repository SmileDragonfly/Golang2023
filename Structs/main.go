package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Contact   string `json:"Contact"`
}

// Request Subscription
type Request_ESubscription struct {
	Mandatory struct {
		TestKey string `json:"TestKey"`
	}
	ProfileID           string `json:"ProfileID"`           //L (1,32) M Merchant’s profile id
	AccessKey           string `json:"AccessKey"`           //L (1,32) M Merchant’s access key
	TransactionID       string `json:"TransactionID"`       //L (1,32) M Identify transaction
	TransactionDateTime string `json:"TransactionDateTime"` //M Transaction date time Format: yyyy-MM-ddTHH:mm:ssZ (UTC date time)
	Language            string `json:"Language"`            //L (2) M Language of SACOMBANK CHECKOUT Page VI: Vietnamese EN: English
	SubscribeOnly       bool   `json:"SubscribeOnly"`       // O Set to true If merchant allows customer to link their
	// cards/accounts without payment or fund transfer.
	// Applied for Sacombank cards/account only
	IsTokenRequest bool `json:"IsTokenRequest"` //M Set True to generate token of card when transaction
	//is processed successfully
	Description string `json:"Description"` // L (1,60) O This value should be something related to customer’s
	//information such as customer id at merchant’s
	//system
	TotalAmount      int64  `json:"TotalAmount"`      // L (1,15) M Transaction’s total amount Set to 0 if SubscribeOnly = true
	DomesticFee      int64  `json:"DomesticFee"`      // L (1,15) O Fee amount applied when customer proceeds transaction at Napas Payment gateway
	InternationalFee int64  `json:"InternationalFee"` //L (1,15) O Fee amount applied when customer proceeds transaction at Sacombank Subscription Gateway.
	SSN              string `json:"SSN"`              // (25) O Customer’s unique id. This value must be whether
	//personal identity number of passport code. Required
	//if SubscribeOnly = true
	Currency  string `json:"Currency"`  //L (3) M The currency of amount. Refer to table Currency
	FirstName string `json:"FirstName"` // L (1,60) M Shipping Address
	//If Country is US and CA, postal code is mandatory
	//Gender: F (female), M (male)
	//Country Code Refer to Country table.
	LastName   string `json:"LastName"`   // L (1,60) M
	Gender     string `json:"Gender"`     //L (1) M
	Address    string `json:"Address"`    //L (1,60) M
	District   string `json:"District"`   //L (1,60) M
	City       string `json:"City"`       //L (1,50) M
	PostalCode string `json:"PostalCode"` //L (1,15) O
	Country    string `json:"Country"`    //M
	Email      string `json:"Email"`      //L (1, 255) MMobile String L (1,15) M
	ReturnUrl  string `json:"ReturnUrl"`  // M The result of authorization transaction will be
	//returned to this URL.
	CancelUrl string `json:"CancelUrl"` // M When customer clicks on Cancel or Back button, web
	//browser will redirect to this URL.
	Signature string `json:"Signature"` // M Signature is calculated in the following steps:
	//1. Concatenate all parameters values except the
	//Signature (arranged in ascending order of
	//parameter’s name) to a string
	//2. RSA sign the hash by merchant’s private key
}

func main() {
	obj := Request_ESubscription{}
	obj.Mandatory.TestKey = "testkey"
	obj.Address = "1"
	obj.AccessKey = "2"
	obj.Email = "3"
	objByte, _ := json.Marshal(obj)
	var o map[string]interface{}
	json.Unmarshal(objByte, &o)
	delete(o, "Signature")
	fmt.Println(o)
	r, _ := json.Marshal(o)
	fmt.Println(string(r))
}
