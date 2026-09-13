package main

import "fmt"

func main() {
	total1, quantos1, oi1 := soma1(10, 10, 1, 2, 3, 5)

	total2, quantos2, oi2 := soma2("tarde", 10, 10, 1, 2, 3, 5)

	fmt.Println(total1, quantos1, oi1)
	fmt.Println(total2, quantos2, oi2)
}

// funções em go podem retornar mais do que um valor,
// pode usar parâmetro variádico
func soma1(x ...int) (int, int, string) {
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma, len(x), "Bom dia!"
}

// em uma função com parâmetro variádico,
// o parâmetro variádico tem que vir no final da declaração da função
func soma2(s string, x ...int) (int, int, string) {
	oi := ""
	switch s {
	case "manhã":
		oi = "Oi, bom dia!"
	case "tarde":
		oi = "Oi, boa tarde!"
	default:
		oi = "Oi, boa noite!"
	}
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma, len(x), oi
}
