package main

import "fmt"

func main() {

	nums := []int{1, 2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum of nums:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("Found 3 at index", i)
		}
	}

	newMap := map[string]string{"banana": "two", "apple": "five"}
	for k, v := range newMap {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range newMap {
		fmt.Println("key:", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c) // Prints index and ASCII value of "g" and "o"
	}

	// range on strings iterates over Unicode code points.
	// The first value is the starting byte index of the rune and the second the rune itself.
	// Rune literals are just 32-bit integer values (represents unicode codepoints)
}
