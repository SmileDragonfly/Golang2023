package main

func main() {
	cards := []string{"Ace of Spades", newCard()}
	println(cards)
}

func newCard() string {
	return "Five of Diamonds"
}
