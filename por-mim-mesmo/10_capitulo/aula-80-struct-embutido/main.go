package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

type profissional struct {
	pessoa
	titulo  string
	salario int
}

func main() {
	pessoa1 := pessoa{
		nome:  "Alfredo",
		idade: 30,
	}

	pessoa2 := profissional{
		pessoa: pessoa{
			nome:  "Maria",
			idade: 25,
		},
		titulo:  "Engenheira",
		salario: 5000,
	}

	fmt.Println(pessoa1)
	fmt.Println(pessoa1.nome)
	fmt.Println(pessoa1.idade)
	fmt.Println()
	fmt.Println(pessoa2)
	fmt.Println(pessoa2.nome)    // Acessando o campo embutido
	fmt.Println(pessoa2.idade)   // Acessando o campo embutido
	fmt.Println(pessoa2.titulo)  // Acessando o campo específico
	fmt.Println(pessoa2.salario) // Acessando o campo específico
	fmt.Println()
}
