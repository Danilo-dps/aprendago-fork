package main

import (
	"fmt"
)

// Bug: depois de close(par) e close(impar), receber de um canal fechado não
// bloqueia — retorna imediatamente o valor zero (0), para sempre. Isso faz
// os cases <-par e <-impar do select ficarem permanentemente prontos, competindo
// com o case <-quit. Como o select escolhe aleatoriamente entre cases prontos,
// o programa pode imprimir várias mensagens erradas antes de encerrar —
// inclusive "O número 0 é ímpar." (via o case impar), mesmo 0 sendo par,
// porque o código não verifica se o valor veio de um canal já fechado.
// Correção: usar o idiom v, ok := <-canal e ignorar quando ok == false,
// ou não depender de receber dos canais par/impar após o fechamento.
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
