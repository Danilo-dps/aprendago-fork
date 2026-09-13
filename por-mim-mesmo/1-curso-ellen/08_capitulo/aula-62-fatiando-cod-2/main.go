package main

import (
	"fmt"
)

func main() {
	//		   			 0.           1.           2.         3.               4.
	sabores := []string{"pepperoni", "mozzarela", "abacaxi", "quatroqueijos", "marguerita"}

	// exclude the elements at index 2 and 3 from the slice
	sabores = append(sabores[:2], sabores[4:]...)

	fmt.Println(sabores)
}
