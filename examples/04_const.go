package main

import (
	"fmt"
	"math"
)

const s string = "constant string"

func main() {
	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)

	//A numeric constant has no type until it’s given one, such as by an explicit cast.
	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
