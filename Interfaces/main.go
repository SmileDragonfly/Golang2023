package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

func (eb englishBot) getGreeting() string {
	return "Hi there!"
}

func (eb englishBot) printGreeting() {
	fmt.Println(eb.getGreeting())
}

func (sb spanishBot) getGreeting() string {
	return "Hola!"
}

func (sb spanishBot) printGreeting() {
	fmt.Println(sb.getGreeting())
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	eb.printGreeting()
	sb.printGreeting()
}
