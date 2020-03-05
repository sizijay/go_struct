package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

type bot interface {
	getGreeting() string
}

func (englishBot) getGreeting() string {
	return "Hello!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	fmt.Print("English bot says: ")
	printGreeting(eb)
	fmt.Println()

	fmt.Print("Spanish bot says: ")
	printGreeting(sb)
	fmt.Println()
}
