package main

import (
	"fmt"
)

// Este exemplo mostra que o BUFFER NÃO ELIMINA O BLOQUEIO — apenas o adia.
//
// make(chan int, 1) cria um canal com espaço para 1 valor.
//   - canal <- 42: cabe no buffer, não bloqueia.
//   - canal <- 43: buffer cheio, BLOQUEIA esperando alguém receber.
//
// Como só existe a goroutine main, e ela está bloqueada no segundo envio,
// ninguém chega para fazer a recepção que liberaria espaço.
// Resultado: deadlock.
//
// Regra: um canal de capacidade N aceita até N envios sem recepção.
// Do N+1 em diante, comporta-se como canal sem buffer — bloqueia.

func main() {
	canal := make(chan int, 1)
	canal <- 42
	canal <- 43
	fmt.Println(<-canal)
}
