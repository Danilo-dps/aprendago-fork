package main

import "fmt"

// mostrando os caracteres da tabela ASCII
func main() {
	for i := 33; i <= 122; i++ {
		fmt.Printf("%d - %v\n", i, string(rune(i)))
	}
}
