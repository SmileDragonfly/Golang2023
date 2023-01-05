package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	Jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			"JimParty@gmail.com",
			94000,
		},
	}
	// fmt.Println(Jim)
	fmt.Printf("%+v", Jim) // value + field name
}
