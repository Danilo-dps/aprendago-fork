package main

import (
	"fmt"
)

func main() {
	//		    0.           1.           2.         3.               4.
	sabores := []string{"pepperoni", "mozzarela", "abacaxi", "quatroqueijos", "marguerita"}

	fatia := sabores[:]

	fmt.Println(fatia)
}
