package main

import (
	"flag"
	"fmt"
	"os"
)

// Greeter is the interface that contains logic of greeting.
type Greeter interface {
	Greet() string
}

// A is an adapter to allow the use of ordinary strings as Greeter.
type A string

// Greet implements Greeter interface.
func (a A) Greet() string {
	return string(a)
}

// B contains logic of repeated greeting.
// It returns its Greeting field repeated Count times.
type B struct {
	Greeting string
	Count    int
}

func (b B) Greet() string {
	var ret string
	for i := 0; i < b.Count; i++ {
		ret += b.Greeting
	}
	return ret
}

// Hello accepts some greeting strategy represented as Greeter interface. It
// then prints it out to the standard output.
func Hello(g Greeter) {
	fmt.Println(g.Greet())
}

// Here declare some type that implements Greeter interface.

func main() {
	var (
		t = flag.String("type", "", "type of greeting")
		n = flag.Int("repeat", 8, "number of greetings for b")
	)
	flag.Parse()

	var g Greeter
	switch *t {
	case "a":
		g = A("hello there!")
	case "b":
		g = B{
			Greeting: "hi!",
			Count:    *n,
		}
	default:
		fmt.Println("no such greeter")
		os.Exit(1)
	}

	Hello(g)
}
