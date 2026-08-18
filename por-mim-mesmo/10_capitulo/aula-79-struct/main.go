package main

import "fmt"

type cliente struct {
	nome      string
	sobrenome string
	fumante   bool
}

func main() {
	cliente1 := cliente{
		nome:      "João",
		sobrenome: "Silva",
		fumante:   false,
	}

	cliente2 := cliente{
		nome:      "Maria",
		sobrenome: "Souza",
		fumante:   true,
	}

	fmt.Println(cliente1)
	fmt.Println(cliente1.nome)
	fmt.Println(cliente1.sobrenome)
	fmt.Println(cliente1.fumante)

	fmt.Println(cliente2)
	fmt.Println(cliente2.nome)
	fmt.Println(cliente2.sobrenome)
	fmt.Println(cliente2.fumante)
}
