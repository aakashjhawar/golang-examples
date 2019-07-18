package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {

	fmt.Println(person{"Steve", 35})             // This creates a new struct
	fmt.Println(person{name: "Franny", age: 23}) // You can name the fields when initializing a struct
	fmt.Println(person{name: "Peter"})           // Omitted fields will be zero-valued
	fmt.Println(&person{name: "Happy", age: 41}) // An & prefix yields a pointer to the struct

	s := person{name: "Parker", age: 17} // Access struct fields with a dot
	fmt.Println(s.name)

	sp := &s // use dots with struct pointers - the pointers are automatically dereferenced
	fmt.Println(sp.age)

	sp.age = 41 // Structs are mutable
	fmt.Println(sp.age)
}
