package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#b53dnd",
	}

	fmt.Println(colors)

	//Way 2:
	var colors2 map[string]string
	fmt.Println(colors2)

	//Way 3:
	colors3 := make(map[string]string)
	//adding to a map
	colors3["white"] = "#gkdkdk"
	fmt.Println(colors3)
	//updating value to a map
	//colors.white = "#34dsfa" --> this will not work
	//always update it like this
	colors3["white"] = "#k9fijf0"
	fmt.Println(colors3)

	//deleting a value from a map
	delete(colors3, "white")
	fmt.Println(colors3)

	//VIDEO 51 ITERATING OVER THE MAP

	colors4 := map[string]string{
		"red":   "#9kdjdn",
		"green": "3343",
		"black": "#ddi4jfd",
	}

	//write a new fuction to iterate over the map and prints out every element on the map
	printMap(colors4)
}

func printMap(c map[string]string) {
	for eachColor, eachColorValue := range c {
		fmt.Println("The Hex code for", eachColor, " is ", eachColorValue)
	}
}

/*
	**VIDEO 49 AND 50
	Both the key and value are statically typed and both key and value can be of different types

	Couple of ways of declaring a map
	Way 1. a way to define a map of key string and value of type string
		map[keyType]valueType

		colors := map[string]string{
			"red":   "#ff0000",
			"green": "#b53dnd",
		}

	Way 2:
		var colors map[string]string // a way to just declare and in this case go will asign default nil value to this map

	Way 3:
		colors3 := make(map[string]string)

	Adding values to a map
		//adding to a map
		colors3["white"] = "#gkdkdk"
		fmt.Println(colors3)
		//updating value to a map
		//colors.white = "#34dsfa" --> this will not work
		//always update it like this
		colors3["white"] = "#k9fijf0"
		fmt.Println(colors3)

		//deleting a value from a map
		delete(colors3, "white")
		fmt.Println(colors3)

	VIDEO 51 : ITERATING OVER THE MAPS

		colors4 := map[string]string{
		"red":   "#9kdjdn",
		"green": "3343",
		"black": "#ddi4jfd",
	}

	//write a new fuction to iterate over the map and prints out every element on the map
		//function definition goes here below
		fnc printMap(argumentName maptype) {
			for mapsKey, mapsValue := range argumentName {

			}
		}

			colors4 := map[string]string{
			"red":   "#9kdjdn",
			"green": "3343",
			"black": "#ddi4jfd",
			}

			//write a new fuction to iterate over the map and prints out every element on the map
			printMap(colors4)

			func printMap(c map[string]string) {
				for eachColor, eachColorValue := range c {
					fmt.Println("The Hex code for", eachColor, " is ", eachColorValue)
				}
			}

	VIDEO 52: DIFFERENCES BETWEEN MAPS AND STRUCTS
			MAP											STRUCT
		1. All keys must be the same type			1. Values can be of different type
		2. Use to represent a collection of 		2. You need to know all the different fields at compile time
			related properties
		3. All values mut be the same type			3. Keys don't support indexing
		4. Don't need to know all the keys 			4. Use to represent a "thing" with a lot of different properites
			at compile time
		5. Keys are indexed - we can				5. Value Type!
			iterate over them
		6. Reference Type! (no need of pointers)


*/
