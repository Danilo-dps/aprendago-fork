package main

import "fmt"

func main() {
	x := 0

	// utilizando o operador & para obter o endereço de x
	y := &x

	fmt.Print(x, y)

	// aqui utilizando o operador * para acessar o valor de x e alterar ele
	*y++

	fmt.Println(*y)
	fmt.Printf("%T, %T\n", x, y)
	fmt.Print(x, y)

	z := &y
	fmt.Printf("\n %v\n", z)
}
