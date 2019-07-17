package main

import "fmt"

func main() {

	var a [7]int
	fmt.Println("array:", a)

	a[5] = 55
	fmt.Println("set the value", a)
	fmt.Println("get the value", a[5])

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	var matrix [3][3]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			matrix[i][j] = i + j
		}
	}

	fmt.Println(matrix)
}
