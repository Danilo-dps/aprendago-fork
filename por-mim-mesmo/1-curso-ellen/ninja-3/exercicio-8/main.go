package main

import "fmt"

func main() {
	x := 4

	switch {
	case x == 5:
		fmt.Println("é igual a 5")
	default:
		fmt.Println("não é igual")
	}
}
