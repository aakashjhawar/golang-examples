package main

import "fmt"

func main() {
	slice := make([]string, 4)
	slice[1] = "hi"
	fmt.Println(slice)

	m := make(map[string]int)
	fmt.Println("empty map:", m)

	m["a"] = 1
	m["b"] = 2
	m["c"] = 3
	m["d"] = 4
	fmt.Println("m", m)
	fmt.Println("lenght of m:", len(m))

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	delete(m, "b")
	fmt.Println("after delete:", m)

	val, prs := m["c"]
	fmt.Println("prs & val:", prs, val)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
