package main

import (
	"fmt"
	"math"
)

const a string = "global string"

func main() {
	fmt.Println(a)

	const b = 5037193
	const c = 4e17 / b
	fmt.Println(b, c)

	fmt.Println(int64(c))
	fmt.Println(math.Sin(b))
}
