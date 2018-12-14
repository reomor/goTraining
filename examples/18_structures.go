package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"Name", 20})
	fmt.Println(person{name: "Alice", age: 20})
	fmt.Println(person{name: "name"})
	fmt.Println(&person{name: "Name", age: 40})

	s := person{name: "Person", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
}
