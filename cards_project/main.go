package main

import (
	"fmt"
)

func main() {

	//cards := []string{"Ace Of Diamons", newCardReturningString()}

	// cards := deck{"Ace Of Diamons", newCardReturningString()}
	// cards = append(cards, "Six of Spades")
	// for i, card := range cards {
	//  fmt.Println(i, card)
	// }
	cards := newDeck()
	cards.print()

	//holding out the two lists returned by the deal method
	hand, remainingDeck := deal(cards, 5)
	hand.print()
	remainingDeck.print()

	// a, b, c := todeals(cards, 5, 6)
	// a.print()
	// b.print()
	// c.print()

	/*
	   Video 25: Type Converstion in go
	   say you wanna convert a string to a byte array
	   byteArrayOfString := []byte("Hi there!")
	       greeting := "Hi there"
	       fmt.Println([]byte(greeting))

	   Question: the deck is a []string and how to convert this array of strings to the byte array
	   firstly we need to convert the array of strings to one solid string and then convert that string to a byte array
	*/

	cardsAsString := cards.toString()
	fmt.Println(cardsAsString)

	//writing the above string into a io write file

	// ioutil.WriteFile("cardsWriteFile.txt", []byte(cardsAsString), fs.ModeAppend)
	cards.writeToAFile("my_cards")

	fmt.Println("priting the cards after reading it from a file")
	myCardsFromDeck := newDeckFromFile("my_cards")
	fmt.Println(myCardsFromDeck)

}

// func newCard() int {
//  return 12
// }

// func newCardReturningString() string {
//  return "Five of Diamons"
// }
