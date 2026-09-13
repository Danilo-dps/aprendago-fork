package main

import (
	"fmt"
)

// Buffer: capacidade de armazenamento do canal.
// make(chan int, 1) cria um canal com buffer de 1 — ou seja, ele já
// nasce com espaço para guardar 1 valor sem que ninguém esteja recebendo.
//
// Efeito: o envio NÃO precisa do encontro simultâneo com a recepção.
// A comunicação deixa de ser síncrona e passa a ser assíncrona.
//
// Aqui só existe a goroutine main. Ela cria o canal com buffer 1,
// envia 42 (cabe no buffer, não bloqueia) e depois recebe. Funciona
// mesmo com uma única goroutine — o que seria impossível sem buffer.
//
// Na prática: canal sem buffer é o padrão. Buffer se usa quando há
// motivo específico (desacoplar ritmos, worker pool, semáforo).

func main() {
	canal := make(chan int, 1)
	canal <- 42
	fmt.Println(<-canal)
}
