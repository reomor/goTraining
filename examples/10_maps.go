package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["k1"] = 1
	m["k2"] = 1111111

	fmt.Println("map is: ", m)

	v1 := m["k1"]
	fmt.Println("v1 is: ", v1)

	fmt.Println("map len is: ", len(m))

	delete(m, "k2")
	fmt.Println("map with deleted by key value: ", m)

	_, prs := m["k2"]
	fmt.Println("present flag: ", prs)

	n := map[string]int{"key1": 1, "key2": 2}
	fmt.Println("static map init: ", n)
}
