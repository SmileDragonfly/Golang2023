package main

import "fmt"

type UserState string

const (
	CREATE UserState = "CREATE"
	ACTIVE UserState = "ACTIVE"
	CLOSED UserState = "CLOSED"
	LOCKED UserState = "LOCKED"
)

func main() {
	var state UserState
	state = "xx"
	fmt.Println("State", state)
}
