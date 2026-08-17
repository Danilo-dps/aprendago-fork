package main

import "fmt"

func main() {
	// inicializando meu slice e atribuindo valores
	meuSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// usando range para percorrer os valores
	for index, value := range meuSlice {
		fmt.Printf("Índice: %d, Valor: %d\n", index, value)
	}

	// mostrando o tipo dele
	fmt.Printf("%T\n", meuSlice)
}
