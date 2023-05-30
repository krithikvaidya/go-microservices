package main

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2*r.Width + 2*r.Height
}

func main() {
	rect := Rectangle{10.0, 15.0}
	fmt.Println(rect.Area())

	// Go automatically handles conversion between values and pointers for method calls. You may
	// want to use a pointer receiver type to avoid copying on method calls or to allow the
	// method to mutate the receiving struct.

	// rp := rect

	fmt.Println("area: ", rect.Area())
	fmt.Println("perim:", rect.Perimeter())
}
