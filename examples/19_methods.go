package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.height * r.width
}

func (r rect) perim() int {
	return 2 * (r.width + r.height)
}

func main() {
	r := rect{width: 10, height: 5}
	fmt.Println("area: ", r.area())
	fmt.Println("perimeter: ", r.perim())

	rPointer := &r
	fmt.Println("area: ", rPointer.area())
	fmt.Println("perim: ", rPointer.perim())
}
