package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) { // Takes an int pointer. The *iptr code in the function body then dereferences
	*iptr = 0 // the pointer from its memory address to the current value at that address.
}

func main() {
	i := 1
	fmt.Println("initial val:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i) // The &i syntax gives the memory address of i, i.e. a pointer to i
	fmt.Println("zeroptr:", i)

}
