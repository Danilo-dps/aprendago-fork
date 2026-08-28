package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

// func (receiver) identifier (parameters) (returns) { code }
func (pr pessoa) euSou() {
	fmt.Printf("Olá, meu nome é %v, meu sobrenome é %v e tenho %v de idade\n", pr.nome, pr.sobrenome, pr.idade)
}

func main() {
	pessoa := pessoa{nome: "Walter", sobrenome: "White", idade: 50}

	pessoa.euSou()
}
