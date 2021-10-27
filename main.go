package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type portugueseBot struct{}

func main() {
	eb := englishBot{}
	pb := portugueseBot{}

	printGreeting(eb)
	printGreeting(pb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	// VERY custom logic for generating an english greeting

	return "Hello, there! how you doing?"
}

func (portugueseBot) getGreeting() string {
	// VERY custom logic for generating a portuguese greeting

	return "E aí!!! como você está?"
}
