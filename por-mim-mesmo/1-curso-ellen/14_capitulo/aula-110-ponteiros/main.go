package main

import (
	"fmt"
)

/*
em go os valores são passados como valor ("pass by value"),
a menos que se passe o endereço do valor.
ao usar o endereço, o valor original é modificado diretamente,
e não uma cópia, o que permite melhor aproveitamento de recursos
ao não precisar ficar copiando os valores.
*/
func main() {
	x := 11

	estarecebeovalor(x)
	fmt.Printf("valor original %v\n", x)

	estarecebeumponteiro(&x)
	fmt.Printf("valor original %v\n", x)
}

// aqui recebe uma cópia, isso não modifica a variavel original
func estarecebeovalor(x int) {
	x++
	fmt.Println("Na função estarecebeovalor:", x)
}

// aqui passasse o endereço, dessa forma acessa o valor atraves do endereço
func estarecebeumponteiro(x *int) {
	*x++
	fmt.Println("Na função estarecebeumponteiro:", *x)
}
