package main

import "fmt"

func main() {
	x := 81

	if x%2 == 0 {
		fmt.Printf("%d é par\n", x)
	}
	if x%2 != 0 {
		fmt.Printf("%d é ímpar\n", x)
	}
	if x%3 == 0 {
		fmt.Printf("%d é divisível por 3\n", x)
	}
	if x%9 == 0 {
		fmt.Printf("%d é divisível por 9\n", x)
	}
	if x%27 == 0 {
		fmt.Printf("%d é divisível por 27\n", x)
	}
}
