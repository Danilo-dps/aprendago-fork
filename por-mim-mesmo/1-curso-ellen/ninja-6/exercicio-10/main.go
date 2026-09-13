package main

import "fmt"

func closure() func() int {
	x := 2
	return func() int {
		x *= x
		return x
	}
}

func main() {
	a := closure()

	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

	b := closure()

	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
}
