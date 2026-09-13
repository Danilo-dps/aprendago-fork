package main

import "fmt"

func main() {
	// inicializando meu array e atribuindo valores
	meuArray := [5]int{1, 2, 3, 4, 5}

	// usando range para percorrer os valores
	for index, value := range meuArray {
		fmt.Println(index, value)
	}

	// mostrando o tipo dele
	fmt.Printf("%T\n", meuArray)
}
