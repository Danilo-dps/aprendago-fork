package main

import (
	"fmt"
	"sync"
)

func main() {
	par := make(chan int)
	impar := make(chan int)
	converge := make(chan int)

	go envia(par, impar)
	go recebe(par, impar, converge)

	for v := range converge {
		fmt.Println("Valor recebido:", v)
	}
}

func envia(p, i chan int) {
	x := 100
	for n := range x {
		if n%2 == 0 {
			p <- n
		} else {
			i <- n
		}
	}
	close(p)
	close(i)
}

func recebe(p, i, c chan int) {
	var wg sync.WaitGroup
	wg.Go(func() {
		for v := range p {
			c <- v
		}
	})
	wg.Go(func() {
		for v := range i {
			c <- v
		}
	})
	wg.Wait()
	close(c)
}

// - Func receive cria duas go funcs, cada uma com um for range, enviando dados dos canais par e ímpar pro canal converge. Não esquecer de WGs!
// - Por fim um range retira todas as informações do canal converge.
