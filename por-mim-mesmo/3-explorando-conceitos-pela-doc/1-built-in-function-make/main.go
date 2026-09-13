package main

import "fmt"

func main() {
	//func make(t Type, size ...IntegerType) Type
	// para slice, primeiro é o tipo, o segundo é a capacidade inicial e o terceiro argumento o quanto ele pode crescer
	slice := make([]int, 5, 10)

	fmt.Println("Valores iniciais do slice:", slice)
	fmt.Println("Tamanho inicial da slice:", len(slice))
	fmt.Println("Capacidade da slice:", cap(slice))
}
