package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

// Um método possui um receiver, que associa o método a um tipo.
// Neste caso, o método oibomdia está associado ao tipo pessoa.
func (p pessoa) oibomdia() {
	fmt.Println(p.nome, "diz bom dia!")
}

// Uma função não possui receiver e, portanto, não pertence
// diretamente a um tipo.
func outraFuncao() {
	fmt.Println("Bom dia!")
}

func main() {
	mauricio := pessoa{"Maurício", 30}

	mauricio.oibomdia()
}
