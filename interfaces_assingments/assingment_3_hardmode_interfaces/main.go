package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Missing parameter, provide file name!")
		return
	}

	fileName := os.Args[1]
	fmt.Println(fileName)

	content, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(content)
	}

	content2, err2 := os.Open(fileName)

	if err2 != nil {
		fmt.Println(err2)
	} else {
		// b := make([]byte, 5) // 5 is the length
		// n, err := f.Read(b)
		// fmt.Printf("%d bytes: %s\n", n, string(b))

		fmt.Println(content2)
	}

	// for _, fileName := range os.Args {
	// 	fmt.Println(fileName)
	// }
	// fmt.Println(len(os.Args), os.Args[1])

}

/*

Create a program that reads the contents of a text file then prints its contents to the terminal.

The file to open should be provided as an argument to the program when it is executed at the terminal. For example
'go run main.go myfile.txt' should open up the  myfile.txt file


To read in the arguments provided to a program, you can reference the variable 'os.Args' which is a slice of type string

To open a file, checkout the documentation for the 'Open' function in the 'os' package - https://golang.org/pkg/os/#Open

What interfaces does the 'File' type implement?

If the 'File' type implements the 'Reader' interface, you might be able to reuse the io.Copy function!


*/
