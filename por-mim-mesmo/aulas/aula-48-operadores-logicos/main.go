package main

import "fmt"

func main() {
	x := 2

	if x == 2 || x == 3 {
		fmt.Println("x é 2 ou 3")
	}

	fmt.Println("===========================")

	y := 6

	if y%2 == 0 && y%3 == 0 {
		fmt.Println("é multiplo de 2 e 3")
	}

	fmt.Println("===========================")

	z := 9

	if (z%2 != 0) && z%3 == 0 {
		fmt.Println("é impar e multiplo de 3")
	}

	fmt.Println("===========================")

	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
	fmt.Println(!false)
}
