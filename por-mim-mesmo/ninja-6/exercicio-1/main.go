package main

import "fmt"

func main() {
	fmt.Println(meuInteiro())
	fmt.Println(umInteiroUmaString())
}

func meuInteiro() int {
	fmt.Println("Minha função de inteiro")
	return 1977
}

func umInteiroUmaString() (int, string) {
	fmt.Println("Minha função de inteiro e string")
	return 1999, "Arroz, feijão, carne e batata"
}
