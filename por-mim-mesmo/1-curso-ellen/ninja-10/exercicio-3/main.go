package main

import (
	"fmt"
)

func main() {
	c := gen()
	receive(c)

	fmt.Println("about to exit")
}

// gen cria um canal bidirecional (chan int) e dispara uma goroutine que envia
// os valores de 0 a 99 nele e depois o fecha (a goroutine usa a variável local c,
// por isso pode enviar e chamar close).
// A assinatura retorna <-chan int: no return c, o canal bidirecional é convertido
// implicitamente para só-recebimento. Assim, quem chama só consegue ler,
// não consegue enviar nem fechar o canal.
func gen() <-chan int {
	c := make(chan int)

	go func() {
		for i := range 100 {
			c <- i
		}
		close(c)
	}()
	return c
}

// recebe como parâmetro um canal que recebe int
func receive(c <-chan int) {
	for i := range c {
		fmt.Println(i)
	}
}

// em um fluxo simples como esse, gen() cria um canal e dispara uma goroutine (via `go func(){...}()`)
// que vai preenchendo esse canal; gen() já retorna `c` imediatamente, sem esperar a goroutine terminar.
// a goroutine main então segue para receive(), que consome o canal aos poucos com `for range c`.
// mesmo que a goroutine disparada por gen() ainda esteja rodando (enviando valores aos poucos),
// receive() consegue ir recebendo cada valor conforme ele chega, sem precisar de nenhuma
// espera manual — quem garante essa sincronização (bloquear até ter dado disponível, ou até fechar)
// é o próprio comportamento do canal.
