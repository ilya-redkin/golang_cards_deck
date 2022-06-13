package main

import "fmt"

func main() {

	fmt.Printf("\n***Before shuffling***\n")
	cards := newDeck()
	cards.print()
	fmt.Printf("\n***After shuffling***\n")
	cards.shuffle()
	cards.print()

}
