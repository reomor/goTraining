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

	l := s[2:5]
	fmt.Println("l = s[2,5] is: ", l)

	l = s[:5]
	fmt.Println("l = s[:5] is: ", l)

	l = s[2:]
	fmt.Println("l = s[2: is: ", l)
}
