package main

import "fmt"

func main() {
	primeiroslice := []int{1, 2, 3, 4, 5}

	fmt.Println(primeiroslice)

	segundoslice := append(primeiroslice[:2], primeiroslice[4:]...)

	fmt.Println(segundoslice)

	fmt.Println(primeiroslice)

	fmt.Println("---------------------------------------------------")

	primeiroSliceOn := []int{1, 2, 3, 4, 5}
	fmt.Println("original:", primeiroSliceOn)

	// primeiroSliceOn[:2:2] usa a sintaxe de 3 índices: [inicio:fim:capacidade]
	// Isso cria uma slice [1,2] com len=2 E cap=2 (em vez de cap=5, que seria
	// o padrão se usássemos só primeiroSliceOn[:2]).
	//
	// Como cap == len, não sobra espaço livre no array subjacente para o
	// append reaproveitar. Isso FORÇA o Go a alocar um array novo e copiar
	// os elementos pra lá, em vez de escrever por cima do array original.
	segundosliceOn := append(primeiroSliceOn[:2:2], primeiroSliceOn[4:]...)

	fmt.Println("novo slice:", segundosliceOn)
	fmt.Println("original (intacto):", primeiroSliceOn)
}
