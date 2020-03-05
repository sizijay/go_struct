package main

import (
	"fmt"
)

func newCard() string {
	//return rand.Intn(100)
	return "Ace of Spades"
}

func main() {
	var T string = "fdhjsdhfjkdhsdjkhjsdfhjksdhfdjksfhsdjkfh"
	fmt.Printf(T)

	card := "Ace of Spades"
	//p := &card
	fmt.Printf("\n %p \n", &card)

	fmt.Println(newCard())
	fmt.Println(newCard())
	fmt.Println()

	cards := newDeck()
	//fmt.Println(cards)

	// for i, num := range cards {
	// 	fmt.Println(i, num)
	// }

	fmt.Println("Calling the print function to print all the values one by one")

	cards.print()
	fmt.Println()
	fmt.Println()

	hand, xDeck := deal(cards, 4)

	fmt.Println("The hand=>")
	hand.print()

	fmt.Println("The remaining deck")
	xDeck.print()
	fmt.Println()
	// fmt.Println(newDeck())

	hand.saveToFile("hand")

	fmt.Println("Opening the deck from the File.....")

	fmt.Println("Correct File Name...")
	A := newDeckFromFile("hand")
	A.print()
	fmt.Println()

	// fmt.Println("Wrong File Name...(Error Handling)")
	// newDeckFromFile("hadsdssdsdnd").print()
	// fmt.Println()

	fmt.Println("Shuffling the cards.....")
	xDeck.shuffle().print()

}
