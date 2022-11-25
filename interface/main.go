package main

import "fmt"

type bot interface {
	getGreeting() string
}
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
func (eb englishBot) getGreeting() string {
	//very custom logic for generating an english greeting
	return "Hello"
}

func (sb spanishBot) getGreeting() string {
	//very custom logic for generating an spanish greeting
	return "Hola!"
}

/*
	VIDEO 53: PURPOSE OF INTERFACES

	We know that
		* Every value has a type
		* Every function has to specify the type of its arguments

	So that mean
		Every function we ever write has to be rewritten to accomodate different
		types evven if the logic in it is identical?

	Problem on a scenario detail:
		Say there are two structs englishBot and spanishBot with two implementations below

		type englishBot struct {
		}
		type spanishBot struct {
		}

		these bots will have two implemtations where one will getGreeting (with custom logic)
		and other to printGreeting (where implementation would be the same). The problem here is that
		we are duplicating the logic of printGreeting when defined both in the struct types

		func(englishBot) getGreeting() string{ return "Hello"}
		func printGreeting(eb englishBot) { fmt.Println(eb.getGreeting())}
		func(spanishBot) getGreeting() string{ return "Hola!"}
		func printGreeting(sb spanishBot) { fmt.Println(sb.getGreeting())}


	VIDEO 54: PROBLEM WITHOUT INTERFACE

	other tip: say oka function undi like below
	func(eb englishBot) getGreeting() string{ return "Hello"}
	//ikkada eb variable vadatle so manam adi tiseyochu to avoid error
		func(englishBot) getGreeting() string {return "Hello"}

		type englishBot struct {}
		type spanishBot struct {}

		these bots will have two implemtations where one will getGreeting (with custom logic)
		and other to printGreeting (where implementation would be the same). The problem here is that
		we are duplicating the logic of printGreeting when defined both in the struct types

		func(englishBot) getGreeting() string{ return "Hello"}
		func printGreeting(eb englishBot) { fmt.Println(eb.getGreeting())}
		func(spanishBot) getGreeting() string{ return "Hola!"}
		func printGreeting(sb spanishBot) { fmt.Println(sb.getGreeting())}

	VIDEO 55: INTERFACES IN PRACTICE
		Using interface were tyring to see if we can overcome the logic duplicaiton over here
		//refactor with interfaces
			type bot interface {
				getGreeting() string
			}
			type englishBot struct{}
			type spanishBot struct{}

			func main() {
				eb := englishBot{}
				sb := spanishBot{}

				printGreeting(eb)
				printGreeting(sb)
			}

			func printGreeting(b bot) {
				fmt.Println(b.getGreeting())
			}
			func (eb englishBot) getGreeting() string {
				//very custom logic for generating an english greeting
				return "Hello"
			}

			func (sb spanishBot) getGreeting() string {
				//very custom logic for generating an spanish greeting
				return "Hola!"
			}

		//explanation: ida epudite oka interface with getGreeting() to defne chestamo
					type bot interface {
					getGreeting() string
					}
		ade type receiver ga petti oka function raste like below
					func (eb englishBot) getGreeting() string {
					//very custom logic for generating an english greeting
					return "Hello"
					}

		epudu ee types (englishBot) anedi bot type kinda maripotundi. apudu ila bot to define
		cheskunna functions anni apply avtayi
		func printGreeting(b bot) {
			fmt.Println(b.getGreeting())
		}

	VIDEO 56 RULES OF INTERFACES
		when you define an interface. the interface will expect any other type to implements
		all of its abstract methods and now when a type implements these methods then the type
		will be considered to be a type of the interface

		//and example of an interface with arguments and errors

		type bot interface {
			//getGreeting(input arguments) (return types)
			getGreeting(string, int) (string, error)
		}

	VIDEO 57 EXTRA INTERFACE NOTES
		1. Interfaces are not generic types --> other languages have 'generic' types where in go does not have any such generics
		2. Interfaces are implicit --> we dont manually say tha our types implement an interface rather they are automatically detected by go where in
			it looks for the methods being implemented and decides on which interface type each type picks on
		3. Interfaces are a contract to help us manage types -> GARBAGE IN -> GARBAGE OUT. If out custom type's implementation of a function is broken then interfaces
																wont help us
		4. Interfaces are touch. Step#1 is understanding how to read them --> understanding how to read interfaces in the standard lib. Writing your own interfaces is tough and
			required experience


	VIDEO 58 THE HTTP PACKAGE
		--> go to interfaces_from58 folder

*/
