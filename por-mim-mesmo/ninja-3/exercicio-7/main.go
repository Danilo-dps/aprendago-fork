package main

import "fmt"

func main() {
	a := 7

	if a%2 == 0 && a%3 == 0 && a%4 == 0 && a%5 == 0 && a%6 == 0 {
		fmt.Println("é um número primo")
	} else {
		fmt.Println("não é um número primo")
	}
}
