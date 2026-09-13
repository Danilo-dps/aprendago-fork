package main

import "fmt"

// fallthrough faz com que a próximo caso seja executado, sem precisar validar a afirmação do booleano em questão
func main() {
	quemtanoescritoriohoje := "ninguem"

	switch quemtanoescritoriohoje {
	case "zezinho":
		fmt.Println("hoje quem tá no escritório é o zezinho")
		fallthrough
	case "marquinhos":
		fmt.Println("hoje quem tá no escritório é o marquinhos")
	case "joana":
		fmt.Println("hoje quem tá no escritório é o joana")
		fallthrough
	case "maria":
		fmt.Println("hoje quem tá no escritório é o maria")
	default:
		fmt.Println("tá vazio. ou a balada tava ótima, ou é feriado")

	}
}
