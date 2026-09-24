package main

import (
	"fmt"
)

// Aqui o `range` itera sobre o canal, e essa iteração só termina quando
// o canal é fechado — canais não fecham sozinhos, é preciso chamar close()
// explicitamente. Sem o close, a goroutine que está recebendo (range) fica
// bloqueada esperando indefinidamente por um novo valor, e como não há mais
// ninguém para enviar, o runtime detecta o deadlock.
func main() {
	c := make(chan int)

	go meuloop(10, c)
	prints(c)
}

func meuloop(t int, s chan<- int) {
	for i := range t {
		s <- i
	}
	close(s)
}

func prints(r <-chan int) {
	for v := range r {
		fmt.Println("Recebido do canal:", v)
	}
}
