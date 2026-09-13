package main

import (
	"fmt"
)

func main() {
	//		   			 0.           1.           2.         3.               4.
	sabores := []string{"pepperoni", "mozzarela", "abacaxi", "quatroqueijos", "marguerita"}

	// how to slice a slice in Go
	// the syntax is slice[start:end] where start is inclusive and end is exclusive
	fatia := sabores[2:4]

	fmt.Println(fatia)
}
