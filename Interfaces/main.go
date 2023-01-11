package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}
type bot interface {
	getGreeting() string
}

func (eb englishBot) getGreeting() string {
	return "Hi there!"
}

func (sb spanishBot) getGreeting() string {
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}
