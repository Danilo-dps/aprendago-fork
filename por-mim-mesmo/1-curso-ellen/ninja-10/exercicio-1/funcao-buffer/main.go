package main

import (
	"fmt"
)

// recapitulando: aqui não há deadlock porque está sendo usado um canal com buffer,
// dessa forma, a mesma goroutine consegue enviar e receber no canal, sincronizando as duas operações.
func main() {
	c := make(chan int, 1)

	c <- 42

	fmt.Println(<-c)
}
