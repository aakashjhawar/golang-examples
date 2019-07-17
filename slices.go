package main

import (
	"fmt"
)

func main() {
	s := make([]string, 3) // To create an empty slice with non-zero length, use the builtin make
	fmt.Println(s)         // Unlike arrays, slices are typed only by the elements they contain (not the number of elements)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set elem:", s)
	fmt.Println("get elem:", s[2])
	fmt.Println("length:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("after append:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("after copy:", c)

	l := c[2:5]
	fmt.Println("split1:", l)
	l = c[:5]
	fmt.Println("split2:", l)
	l = c[2:]
	fmt.Println("split3:", l)

	t := [5]string{"a", "b", "c", "d", "e"} // declare and initialize
	fmt.Println("new array:", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)
}
