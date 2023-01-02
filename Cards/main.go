package main

import "fmt"

func main() {
	cards := newDeck()
	hand, remainingCards := deal(cards, 5)
	fmt.Println("All cards:")
	cards.print()
	fmt.Println("Hand:")
	hand.print()
	fmt.Println("Remaining cards:")
	remainingCards.print()
}
