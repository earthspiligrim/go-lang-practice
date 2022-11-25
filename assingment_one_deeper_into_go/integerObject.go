package main

import "fmt"

type intarray []int

func integerToReturn() intarray {
	return []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
}

func (inputIntArray intarray) evenOrOdd() {
	for _, eachnumber := range inputIntArray {
		if eachnumber%2 == 0 {
			fmt.Println("number is even")
		} else {
			fmt.Println("number is odd")
		}
	}
}
