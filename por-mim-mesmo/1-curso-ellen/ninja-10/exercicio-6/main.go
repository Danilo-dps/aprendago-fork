package main

import "fmt"

func main() {
	c := make(chan int)

	go meuloop(100, c)
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
