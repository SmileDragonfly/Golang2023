package main

func main() {
	cards := []string{"Ace of Spades", newCard()}
	for i, card := range cards {
		println(i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
