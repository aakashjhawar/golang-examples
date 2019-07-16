package main

import (
	"fmt"
	"math"
)

const a string = "global string"

func main() {
	fmt.Println(a)

	const b = 500000000
	const c = 3e20 / b
	fmt.Println(b, c)

	fmt.Println(int64(c))
	fmt.Println(math.Sin(b))
}
