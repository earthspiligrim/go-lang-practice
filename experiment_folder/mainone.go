package main

import "fmt"

func main() {
	//printState()

	m := make(map[int]string)

	m[1] = "a"
	m[2] = "b"
	m[3] = "c"
	m[4] = "d"

	// Can this be done better?
	v := make([]string, len(m), len(m))
	idx := 0
	for _, value := range m {
		v[idx] = value
		idx++
	}

	fmt.Println(v)

	// declaring an array

	// var players[4] string {"jkk", "ddd", "dadd", "ddad"}
	arr := [4]string{"kkk", "dksakd", "adfd", "dadd"}

	for _, eacharr := range arr {
		fmt.Println(eacharr)
	}
}
