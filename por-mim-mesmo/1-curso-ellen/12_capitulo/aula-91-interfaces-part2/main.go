package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

type arquiteto struct {
	pessoa
	especialização  string
	obrasconcluidas int
}

type dentista struct {
	pessoa
	especialização   string
	dentesarrancados int
}

func (a arquiteto) oibomdia() {
	fmt.Println("Oi, bom dia! Meu nome é", a.nome, "e eu já fiz", a.obrasconcluidas, "prédios.")
}

func (d dentista) oibomdia() {
	fmt.Println("Oi, bom dia! Meu nome é", d.nome, "e eu já arranquei", d.dentesarrancados, "dentes.")
}

type gente interface {
	oibomdia()
}

func serhumano(g gente) {
	switch g := g.(type) {
	case arquiteto:
		fmt.Println("Este arquiteto é especializado em", g.especialização, "e já fez", g.obrasconcluidas, "obras. Ele diz:")
	case dentista:
		fmt.Println("Este dentista é especializado em", g.especialização, "e já arrancou", g.dentesarrancados, "dentes. Eles diz:")
	}

	g.oibomdia()
}

func main() {
	pessoa1 := arquiteto{
		pessoa: pessoa{
			nome:      "Paulo",
			sobrenome: "Prédio",
			idade:     40,
		},
		especialização:  "Galpão de fazenda",
		obrasconcluidas: 20,
	}

	pessoa2 := dentista{
		pessoa: pessoa{
			nome:      "Henrique",
			sobrenome: "Cido",
			idade:     50,
		},
		especialização:   "Tortura",
		dentesarrancados: 8748,
	}

	// fmt.Println(pessoa1, pessoa2)

	//pessoa1.oibomdia()
	//pessoa2.oibomdia()

	serhumano(pessoa1)
	serhumano(pessoa2)
}
