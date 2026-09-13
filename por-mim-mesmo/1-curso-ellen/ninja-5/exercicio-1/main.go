package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	sabores   []string
}

func main() {
	pessoa1 := pessoa{
		nome:      "João",
		sobrenome: "Silva",
		sabores:   []string{"Chocolate", "Morango", "Baunilha"},
	}

	pessoa2 := pessoa{
		nome:      "Maria",
		sobrenome: "Silva",
		sabores:   []string{"Flocos", "Milho verde", "Doce de leite", "Brigadeiro"},
	}

	fmt.Println(pessoa1.nome, pessoa1.sobrenome)

	// O range em slices sempre retorna dois valores: (índice, conteúdo).
	// Assim como funções em Go podem retornar múltiplos valores,
	// o underscore aqui serve apenas para descartar o índice,
	// que não será utilizado, e manter só o conteúdo em "sabor".
	for _, sabor := range pessoa1.sabores {
		fmt.Println(sabor)
	}

	fmt.Println("===================================")

	fmt.Println(pessoa2.nome, pessoa2.sobrenome)
	for _, sabor := range pessoa2.sabores {
		fmt.Println(sabor)
	}
}
