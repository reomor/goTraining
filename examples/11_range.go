package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4}
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	fmt.Println("sum is: ", sum)

	for i, num := range numbers {
		if num == 3 {
			fmt.Println("index is: ", i)
		}
	}

	kvs := map[string]string{"a": "aaa", "b": "bbb"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key is: ", k)
	}

	for i, ch := range "go go go" {
		fmt.Println(i, ch)
	}
}
