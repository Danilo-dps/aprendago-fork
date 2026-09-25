package main

import (
	"fmt"
)

//        - Chans par, ímpar, quit
//        - Func send manda números pares pra um canal, ímpares pra outro, e fecha/quit
//        - Func receive é um select entre os três canais, encerra no quit

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
		case v := <-par:
			fmt.Println("O número", v, "é par.")
		case v := <-impar:
			fmt.Println("O número", v, "é ímpar.")
		case <-quit:
			return
		}
	}
}
