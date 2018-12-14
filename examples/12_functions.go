package main

import "fmt"

func plusDuo(a int, b int) int {
	return a + b
}

func plusTriple(a, b, c int) int {
	return a + b + c
}

func main() {
	res := plusDuo(3, 7)
	fmt.Println("5 + 7 = ", res)
	res = plusTriple(3, 5, 7)
	fmt.Println("3 + 5 + 7 = ", res)
}
