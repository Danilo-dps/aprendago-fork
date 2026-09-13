package main

import "fmt"

func main() {
	// inicializando meu slice e atribuindo valores
	meuSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// O operador '...' após o slice "espalha" seus elementos como argumentos
	// individuais para a função variádica soma(x ...int). Sem o '...', o Go
	// não aceitaria passar um []int diretamente onde se espera argumentos
	// separados — é o inverso do empacotamento que a própria função faz
	// ao juntar os argumentos recebidos de volta em um slice (x).
	total1 := somaVariadica(meuSlice...)

	total2 := somaSlice(meuSlice)

	fmt.Println(total1)
	fmt.Println(total2)
}

// soma recebe um número variável de argumentos int (parâmetro variádico).
// Internamente, o Go empacota todos os valores passados na chamada em um
// único slice ([]int), que é o que 'x' realmente é dentro da função.
//
// É por isso que, ao chamar soma com um slice já existente, é preciso usar
// o operador '...' (ex: soma(meuSlice...)): ele desempacota o slice em
// argumentos individuais, que é o formato que essa assinatura espera.
// Sem o '...', passar um []int diretamente não compila, pois um slice
// (uma unidade só) não é automaticamente equivalente a "vários int soltos".
func somaVariadica(x ...int) int {
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma
}

// Aqui, como a semântica diz, a função espera UM slice
func somaSlice(slice []int) int {
	soma := 0
	for _, v := range slice {
		soma += v
	}
	return soma
}
