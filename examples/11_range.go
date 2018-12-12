package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4}
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	fmt.Println("sum is: ", sum)
}
