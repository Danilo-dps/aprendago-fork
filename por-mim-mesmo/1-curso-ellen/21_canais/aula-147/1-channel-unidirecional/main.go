package main

import (
	"fmt"
)

func main() {
	// aqui é criado um canal que pode enviar e receber valores do tipo int
	// ainda sem definir se é send ou receive
	canal := make(chan int)

	go send(canal)

	receive(canal)
}

// função que envia um valor para o canal
// note que a sintaxe é chan<- int, indicando que o canal apenas envia dados(coloca dado dentro do canal), não serve para receber
func send(s chan<- int) {
	s <- 42
}

// função que recebe um valor do canal
// note que a sintaxe é <-chan int, indicando que o canal apenas recebe dados(pega dado do canal), não serve para enviar
func receive(r <-chan int) {
	fmt.Println("O valor recebido do canal foi:", <-r)
}
