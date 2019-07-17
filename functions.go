package main

import (
	"fmt"
)

func addTwoNumber(a int, b int) int {
	return a + b
}

func addThreeNumber(a, b, c int) int {
	return a + b + c
}

func vals() (int, int) {
	return 3, 7
}

func main() {
	fmt.Println("Main Function")

	res := addTwoNumber(2, 3)
	fmt.Println("2 + 3 =", res)

	res = addThreeNumber(4, 5, 7)
	fmt.Println("4 + 5 + 7 =", res)

	a, b := vals()
	fmt.Println("a & b:", a, b)
}
