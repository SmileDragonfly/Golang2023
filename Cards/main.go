package main

func main() {
	cards := deck{"Ace of Spades", newCard()}
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
