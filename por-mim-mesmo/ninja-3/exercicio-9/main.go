package main

import "fmt"

func main() {
	esporteFavorito := "Futebol"

	switch esporteFavorito {
	case "Futebol":
		fmt.Println("Futebol é o meu esporte favorito")
	case "Volei":
		fmt.Println("Volei é o meu esporte favorito")
	case "Basquete":
		fmt.Println("Basquete é o meu esporte favorito")
	case "Skatismo":
		fmt.Println("Skatismo é o meu esporte favorito")
	default:
		fmt.Println("nenhum esporte é o favorito")
	}
}
