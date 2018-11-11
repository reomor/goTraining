package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("empty a: ", a)

	a[4] = 104
	fmt.Println("a is: ", a)
	fmt.Println("a[4] is: ", a[4])

	fmt.Println("a length is:", len(a))

	b := [5]int{1, 2, 3}
	fmt.Println("b is: ", b)

	var twoDimensional [2][3]int
	for i := 0; i < len(twoDimensional); i++ {
		for j := 0; j < len(twoDimensional[0]); j++ {
			twoDimensional[i][j] = 2*i + j + i
		}
	}
	fmt.Println(twoDimensional)
}
