package main

import "fmt"

func main() {
	umaslice := []int{1, 2, 3, 4}
	outraslice := []int{9, 10, 11, 12}

	fmt.Println(umaslice)

	umaslice = append(umaslice, 5, 6, 7, 8)

	fmt.Println(umaslice)

	// slice de int é um tipo de slice, então podemos anexar um slice a outro slice usando o operador "..." (spread operator)
	// o operador "..." é usado para expandir o slice em seus elementos individuais, permitindo que eles sejam adicionados ao outro slice
	// slice tem como tipo principal ser uma slice, por isso não podemos anexar um slice a outro slice diretamente,
	// mas podemos usar o operador "..." para expandir o slice em seus elementos individuais e adicioná-los ao outro slice
	umaslice = append(umaslice, outraslice...)

	fmt.Println(umaslice)
}
