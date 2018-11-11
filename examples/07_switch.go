package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("it's after noon")
	}

	whatIAm := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("i'm a bool")
		case int:
			fmt.Println("i'm a int")
		default:
			fmt.Printf("i don't know who i am :( '%T'\n", t)
		}
	}

	whatIAm(true)
	whatIAm(45)
	whatIAm("ololo")
}
