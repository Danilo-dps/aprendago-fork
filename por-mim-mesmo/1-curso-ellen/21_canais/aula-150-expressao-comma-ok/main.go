package main

import (
	"fmt"
)

func main() {
	par := make(chan int)
	impar := make(chan int)
	quit := make(chan bool)

	go mandaNumeros(par, impar, quit)

	receive(par, impar, quit)
}

func mandaNumeros(par, impar chan int, quit chan bool) {
	total := 100
	for i := range total {
		if i%2 == 0 {
			par <- i
		} else {
			impar <- i
		}
	}
	close(par)
	close(impar)
	quit <- true
}

func receive(par, impar chan int, quit chan bool) {
	for {
		select {
		case v, ok := <-par:
			if !ok {
				par = nil // desativa esse case no select
				continue
			}
			fmt.Println("O número", v, "é par.")
		case v, ok := <-impar:
			if !ok {
				impar = nil // desativa esse case no select
				continue
			}
			fmt.Println("O número", v, "é ímpar.")
		case <-quit:
			fmt.Println("Encerrando.")
			return
		}
	}
}
