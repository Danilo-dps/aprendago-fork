package main

import "fmt"

func main() {
	// defer é uma instrução que agenda a execução de uma função
	// para quando a função que contém o defer estiver prestes a retornar.
	//
	// Quando existem vários defer, eles são executados na ordem inversa
	// em que foram registrados: o último defer é executado primeiro.
	defer fmt.Println("defer é um statement que agenda uma função para ser executada no retorno da função")
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	fmt.Println(4)
}
