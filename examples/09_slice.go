package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("empty s is: ->", s, "<-")
	fmt.Println("s len is: ", len(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set: ", s)
	fmt.Println("get [2]: ", s[2])
	fmt.Println("size s is: ", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("s is: ", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy is: ", c)

	l := s[2:5]
	fmt.Println("l = s[2,5] is: ", l)

	l = s[:5]
	fmt.Println("l = s[:5] is: ", l)

	l = s[2:]
	fmt.Println("l = s[2: is: ", l)

	t := []string{"x", "y", "z"}
	fmt.Println("static init: ", t)

	twoD := make([][]int, 3)
	for i := 0; i < len(twoD); i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j + 1
		}
	}
	fmt.Println("2d arr: ", twoD)
}
