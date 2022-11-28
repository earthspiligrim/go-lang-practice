package main

import "fmt"

type shape interface {
	findArea() float64
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

func (sq square) findArea() float64 {
	return sq.sideLength * sq.sideLength
}

func (tq triangle) findArea() float64 {
	return 0.5 * tq.height * tq.height
}

func printArea(sp shape) {
	fmt.Println(sp.findArea())
}

func main() {
	tri := triangle{
		base:   43.3,
		height: 23.90,
	}

	printArea(tri)

	sq := square{sideLength: 433}
	printArea(sq)
}

/*
	Write a program that creates two custom struct types called 'triangle' and 'square'

	The 'square' type should be a struct with a field called 'sideLength' of type float64

	The 'triangle' type should be a struct with a field called 'height' of type float64 and a field of type 'base' of type float64

	Both types should have function called 'getArea' that returns the calculated area of the square or triangle

	Area of a triangle = 0.5 * base * height.
	Area of a square = sideLength * sideLength

	Add a 'shape' interface that defines a function called 'printArea'. This function should calculate the area of the given shape and
	print it out to the terminal Design the interface so that the 'printArea' function can be called with either a triangle or a square.



*/
