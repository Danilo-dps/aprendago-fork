package main

import "fmt"

func main() {
	a := make(chan int)
	b := make(chan int)
	x := 500

	// aqui a goroutine envia dados para o canal a
	go func(x int) {
		for i := range x {
			a <- i
		}
	}(x / 2)

	// aqui a goroutine envia dados para o canal b
	go func(x int) {
		for i := range x {
			b <- i
		}
	}(x / 2)

	// aqui é feito um loop limitado pelo valor de x,
	// e dentro do loop é feito um select para receber dados dos canais a e b.
	// select é o equivalente a um switch, mas para canais: ele bloqueia até que
	// uma das operações de envio ou recebimento esteja pronta.
	// se mais de um case estiver pronto ao mesmo tempo, o Go escolhe um deles
	// aleatoriamente, por isso não há garantia de ordem entre os valores
	// recebidos do canal A e do canal B na saída.
	for range x {
		select {
		case v := <-a:
			fmt.Println("Canal A:", v)
		case v := <-b:
			fmt.Println("Canal B:", v)
		}
	}
}
