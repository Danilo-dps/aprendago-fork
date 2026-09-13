package main

import "fmt"

func main() {
	// inicializando meu slice e atribuindo valores
	meuSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(meuSlice[:3])
	fmt.Println(meuSlice[4:])
	fmt.Println(meuSlice[1:7])
	fmt.Println(meuSlice[2:9])
	fmt.Println(meuSlice[2 : len(meuSlice)-1])
}
