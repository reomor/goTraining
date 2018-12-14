package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(ival *int) {
	*ival = 0
}

func main() {
	i := 1
	fmt.Println("initial value is: ", i)

	zeroval(i)
	fmt.Println("zeroval execution: ", i)
	zeroptr(&i)
	fmt.Println("zeroptr execution: ", i)

	fmt.Println("pointer is: ", &i)
}
