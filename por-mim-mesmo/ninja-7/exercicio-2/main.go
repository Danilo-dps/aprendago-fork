package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

// ao trabalhar com ponteiro para struct, o Go permite acessar os campos
// diretamente (ex: p.idade) sem precisar desreferenciar explicitamente (*p).idade
// isso só existe para structs (e chamada de métodos)
// por baixo dos panos, o compilador faz a desreferência automaticamente
func mudeMe(p *pessoa) {
	(*p).nome = "João"
	(*p).sobrenome = "Pereira"
	p.idade = 30 // forma implícita de desreferenciar, válida por 'p' apontar para um struct
}

func main() {
	pe := pessoa{nome: "Lucas", sobrenome: "Silva", idade: 29}

	fmt.Println(pe.nome, pe.sobrenome, pe.idade)
	mudeMe(&pe)
	fmt.Println(pe.nome, pe.sobrenome, pe.idade)
}
