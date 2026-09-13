package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

// func (receiver) identifier (parameters) (returns) { code }
func (p *pessoa) falar() {
	fmt.Printf("Olá, meu nome é %v, meu sobrenome é %v e tenho %v anos\n", p.nome, p.sobrenome, p.idade)
}

type humano interface {
	falar()
}

func dizerAlgumaCoisa(h humano) {
	h.falar()
}

func main() {
	mp := pessoa{"Jhon", "Snow", 25}

	dizerAlgumaCoisa(&mp)
}

// nesse exemplo a função dizerAlgumaCoisa só funciona com &mp
// porque o método falar foi definido com receiver *pessoa (ponteiro)
// em Go, quando um método tem receiver de ponteiro, é o tipo *pessoa que implementa a interface, e não o tipo pessoa (valor)
// por isso, para passar mp como argumento de tipo humano, é obrigatório passar o ponteiro &mp
//
// regra geral em Go:
// - se o método tem receiver *T, apenas *T implementa a interface
// - se o método tem receiver T, tanto T quanto *T implementam a interface
//   (porque Go consegue tirar o endereço automaticamente quando você tem uma variável endereçável)
