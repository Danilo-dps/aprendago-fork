package main

import "fmt"

func main() {
	meuMap := map[string][]string{
		"einstein albert": {
			"imaginar", "refletir", "se questionar",
		},
		"curie marie": {
			"descobrir a radiação", "ser uma mulher incrível", "ganhar dois prêmios nobel",
		},
		"ritchie dennis": {
			"criar a linguagem de programação C", "criar o sistema operacional UNIX", "inspirar milhões de programadores",
		},
	}

	meuMap["lovelace ada"] = []string{"escrever o primeiro programa de computador", "gostar de lógica", "inspirar milhões de programadores"}

	for k, v := range meuMap {
		fmt.Println(k)
		for i, h := range v {
			fmt.Println("\t", i, " - ", h)
		}
	}
}
