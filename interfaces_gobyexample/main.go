package main

import "fmt"

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	//width, height float64
	r := rectangle{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}
