package main

import "fmt"

func main() {
	cards := deck{"Ace of Spades", newCard()}
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
