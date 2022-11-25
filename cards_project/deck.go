package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

//creat a new type of 'deck'
//whhich is a slice of strings

type deck []string
type intDeck []int
type someDeck string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "four"}

	// to tell go that you dont wanna use a variable yet u need to declare it like below
	// kind regular for loop lo index define pakka cheyali and vadakpote compiler tidtadi
	// apudu _ pettinam ante go ki ignore cheyi nenu vada but syntax prakaram nenu vadali ani
	//cheptunattu

	for _, eachCardSuit := range cardSuits {
		for _, eachCardValue := range cardValues {
			cards = append(cards, eachCardValue+" of "+eachCardSuit)
		}
	}
	return cards
}

func (d deck) print() {
	for index, eachCard := range d {
		fmt.Println(index, eachCard)
	}
}

// to return two values from the function you can say (returntype1, returntype2)
func deal(d deck, handSize int) (deck, deck) {
	//slicing a slice based on a range
	/*
	   say you have a slice arr := [] int {3,5,6,7,5}
	   slicing the slice would be like arr[startindex, endindex]
	   the about slice would include the element at the start index
	   yet exclude the element at the endindex

	   when you write arr[startindex:] you wanna choose the tail of the list including the start index
	   when you write arr[:endindex] you wanna choose the head of the list starting first element till element at end index(exclusive)
	*/
	return d[:handSize], d[handSize:]
}

// func todeals(d deck, handSizeOne int, handSizeTwo int) (deck, deck, deck) {
//  return d[:handSizeOne], d[handSizeOne+1 : handSizeTwo], d[handSizeTwo:]
// }

//packages are the go package which is used

func (d deck) toString() string {
	// var resultString string = ""
	// for _, eachString := range d {
	//  resultString += eachString + ", "
	// }
	// return resultString

	/*
	   there is a helper package (strings) can deal with strings
	   the join function takes a slice of string and joins together with a sep string and return as a string
	   func Join(a [] string, sep string) string
	*/

	return strings.Join([]string(d), ",")
}

func (d deck) writeToAFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	/*
	   ioutil.ReadFile(fileName string) ([]byte, error)
	*/
	bytesOfCards, err := ioutil.ReadFile(filename)

	if err != nil {
		// strings.
		fmt.Println("Error:", err)
		//using the os package of go we can exit the program entirely
		//os.Exit(code int) the integer 0 -> all good and other than 0 indicates
		//that something has gone wrong
		os.Exit(1)
	}

	s := strings.Split(string(bytesOfCards), ",")
	return deck(s)
}

// shuffle fucntion
// to get the length of the slice len(d)
// func (d deck) shuffleCards() {
//  for index := range d {
//      newPosition := rand.Intn(len(d) - 1)
//      //the above random generator is a sudo random generator
//      // go uses the input as a aseed and

//      d[index], d[newPosition] = d[newPosition], d[index]
//  }
// }

// writing the random numbers in a different way
// Revisit Video 31
func (d deck) shuffleCards() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for index := range d {
		newPosition := r.Intn(len(d) - 1)
		d[index], d[newPosition] = d[newPosition], d[index]
	}
}
