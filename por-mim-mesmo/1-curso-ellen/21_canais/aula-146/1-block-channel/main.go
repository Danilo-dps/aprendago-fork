package main

import (
	"fmt"
)

// Canais são o mecanismo de comunicação entre goroutines em Go.
// A filosofia: "não comunique compartilhando memória; compartilhe memória se comunicando".
// Criados com make(chan T).
//
// Canal SEM buffer sincroniza envio e recepção: eles precisam se encontrar
// no mesmo instante (handshake), como passar um bastão — A chega, B pega.
//
// A operação de ENVIO (linha 21, canal <- 42) bloqueia até que OUTRA
// goroutine execute a RECEPÇÃO (<-canal). Como aqui só existe a goroutine main
// e ela está bloqueada no envio, ninguém chega para receber: deadlock
// O runtime detecta que todas as goroutines dormem e aborta.

func main() {
	canal := make(chan int)
	canal <- 42
	fmt.Println(<-canal)
}
