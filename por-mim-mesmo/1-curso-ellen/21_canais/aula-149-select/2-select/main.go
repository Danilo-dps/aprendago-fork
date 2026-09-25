package main

import (
	"fmt"
)

func main() {
	canal := make(chan int)
	quit := make(chan int)
	go recebeQuit(canal, quit)
	enviaPraCanal(canal, quit)
}

// aqui o loop roda 50 vezes (for range 50, sobre o inteiro, não sobre o canal),
// e a cada iteração recebe um valor do canal com <-canal.
// depois de receber os 50 valores, envia um valor para quit,
// avisando enviaPraCanal que deve parar.
func recebeQuit(canal chan int, quit chan int) {
	for range 50 {
		fmt.Println("Recebido:", <-canal)
	}
	quit <- 0
}

// quem está colocando dado no canal é essa função aqui, conforme é colocado no select,
// ele vai colocar o valor de qualquercoisa no canal, e depois incrementa.
// Se o quit receber algum valor, ele retorna e encerra a função.
func enviaPraCanal(canal chan int, quit chan int) {
	qualquercoisa := 1
	for {
		select {
		case canal <- qualquercoisa:
			qualquercoisa++
		case <-quit:
			return
		}
	}
}

// diferente de fechar o canal com close(), aqui NÃO há fechamento de canal.
// a sincronização é feita por um canal de sinalização (quit): recebeQuit conta
// 50 recebimentos e então ENVIA um valor em quit (não fecha), avisando
// enviaPraCanal que deve parar. Isso ocorre de forma concorrente:
// enviaPraCanal fica preenchendo canal, recebeQuit fica consumindo,
// e quando termina, sinaliza via quit para enviaPraCanal retornar.
