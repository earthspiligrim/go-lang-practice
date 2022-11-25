package main

type rectangle struct {
	width, height float64
}

func (rect rectangle) area() float64 {
	return rect.height * rect.width
}

func (rect rectangle) perim() float64 {
	return (2 * rect.height) + (2 * rect.width)
}
