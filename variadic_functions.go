package main

import (
	"fmt"
)

// Variadic functions can be called with any number of trailing arguments. Eg, Println

func sum(nums ...int) {
	fmt.Println(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println("total", total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 4, 5)

	nums := []int{11, 12, 13, 14}
	sum(nums...)
}
