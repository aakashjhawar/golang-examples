package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int { // This area method has a receiver type of *rect
	return r.width * r.height
}

func (r rect) peri() int { // Methods can be defined for either pointer or value receiver types
	return 2 * (r.width + r.height)
}

func main() {
	r := rect{width: 10, height: 30}

	fmt.Println("Area:", r.area())
	fmt.Println("Permieter:", r.peri())

	rp := &r
	rp.height = 5
	rp.width = 3
	fmt.Println("Area: ", rp.area())
	fmt.Println("Perimeter:", rp.peri())

	fmt.Println("r:", r)
	fmt.Println("rp:", rp)

	// Go automatically handles conversion between values and pointers for method calls
	// Use a pointer receiver type to avoid copying on method calls or to allow the method to mutate the receiving struct

}
