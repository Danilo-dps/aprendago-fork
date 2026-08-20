package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	sabores   []string
}

func main() {
	// make() aloca e inicializa a estrutura interna do map, retornando
	// um map vazio, porém utilizável. Sem o make (ou um literal como
	// map[string]pessoa{}), a variável teria valor zero "nil", e
	// tentar escrever nela (mepe["chave"] = valor) causaria um panic
	// em tempo de execução: "assignment to entry in nil map".
	// Leitura em um map nil é permitida (retorna o valor zero do tipo),
	// mas escrita não.
	mepe := make(map[string]pessoa)

	mepe["Silva"] = pessoa{
		nome:      "João",
		sobrenome: "Silva",
		sabores:   []string{"Chocolate", "Morango", "Baunilha"},
	}

	mepe["Pereira"] = pessoa{
		nome:      "Maria",
		sobrenome: "Pereira",
		sabores:   []string{"Flocos", "Milho verde", "Doce de leite", "Brigadeiro"},
	}

	for _, valor := range mepe {
		fmt.Println("Meu nome é", valor.nome, "e meus sorvetes favoritos são:")
		for _, valor := range valor.sabores {
			fmt.Println("–", valor)
		}
		fmt.Println()
	}

	fmt.Println("-------------------------------------------------")

	// Aqui não é necessário usar make(), porque o map já está sendo
	// criado e populado numa única expressão: um "literal composto"
	// (map[string]pessoa{...}). Diferente do mepe acima, que nasce
	// vazio e é preenchido depois via atribuição (mepe["chave"] = ...),
	// aqui o Go aloca o map e já insere todos os pares chave-valor
	// no mesmo passo. O make só é necessário quando você precisa de
	// um map vazio para popular em momentos separados do código.
	mepe2 := map[string]pessoa{
		"Pimentão": {
			nome:      "Renata",
			sobrenome: "Pimentão",
			sabores:   []string{"pistache", "morango", "baunilha"},
		},
		"da Prússia": {
			nome:      "Frederico",
			sobrenome: "da Prússia",
			sabores:   []string{"sabão em pó", "pé de coelho", "feijão"},
		},
	}

	for _, valor := range mepe2 {
		fmt.Println("Meu nome é", valor.nome, "e meus sorvetes favoritos são:")
		for _, valor := range valor.sabores {
			fmt.Println("–", valor)
		}
		fmt.Println()
	}
}
