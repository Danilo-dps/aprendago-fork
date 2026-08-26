package main

import "fmt"

func main() {
	si := []int{10, 10, 1, 2, 3, 5}

	// passando o operador ... para que a função consiga operar
	total1 := soma(si...)

	// uma função variádica funciona se não passar nenhum argumento para ela
	total2 := soma()

	fmt.Println(total1)
	fmt.Println(total2)
}

// funções em go podem retornar mais do que um valor,
// pode usar parâmetro variádico
func soma(x ...int) int {
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma
}
