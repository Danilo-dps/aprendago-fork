package main

import (
	"fmt"
)

// Não há deadlock porque agora existem DUAS goroutines.
//
// A goroutine main:
//  1. cria o canal (make)
//  2. dispara a função anônima com `go` (não espera ela rodar)
//  3. bloqueia em <-canal, esperando receber
//
// A goroutine anônima:
//   - envia 42 para o canal, bloqueando até a main receber
//
// O canal sincroniza as duas: envio e recepção se encontram e ambas
// destravam. A ORDEM em que cada goroutine chega ao encontro é
// indeterminada, mas isso não importa — o canal garante o handshake.

func main() {
	canal := make(chan int)

	go func() {
		canal <- 42
	}()

	fmt.Println(<-canal)
}
