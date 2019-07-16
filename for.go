package main

import (
	"fmt"
)

func main() {
	i := 1
	for i <= 2 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 1; j < 17; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for k := 0; k <= 8; k++ {
		if k%2 == 0 {
			continue
		}
		fmt.Println(k)
	}
}
