package main

import (
	"fmt"
)

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

type dentista struct {
	pessoa
	dentesarrancados int
	salário          float64
}

type arquiteto struct {
	pessoa
	tipodeconstrução string
	tamanhodaloucura string
}

func (x dentista) oibomdia() {
	fmt.Println("Meu nome é", x.nome, "e eu já arranquei", x.dentesarrancados, "dentes, e ouve só: Bom dia!")
}

func (x arquiteto) oibomdia() {
	fmt.Println("Meu nome é", x.nome, "e ouve só: Bom dia!")
}

type gente interface {
	oibomdia()
}

func serhumano(g gente) {
	g.oibomdia()
	switch g := g.(type) {
	case dentista:
		fmt.Println("Eu ganho:", g.salário)

	case arquiteto:
		fmt.Println("Eu construo:", g.tipodeconstrução)
	}
}

func main() {
	mrdente := dentista{
		pessoa: pessoa{
			nome:      "Mister Dente",
			sobrenome: "da Silva",
			idade:     50,
		},
		dentesarrancados: 8973,
		salário:          98736.06,
	}

	mrprédio := arquiteto{
		pessoa: pessoa{
			nome:      "Mister Prédio",
			sobrenome: "Sobrenome",
			idade:     51,
		},
		tipodeconstrução: "Hospícios",
		tamanhodaloucura: "Demais",
	}

	serhumano(mrdente)
	fmt.Println("")
	serhumano(mrprédio)
}
