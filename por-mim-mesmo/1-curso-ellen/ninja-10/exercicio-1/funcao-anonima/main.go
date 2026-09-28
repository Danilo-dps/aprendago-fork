package main

import (
	"fmt"
)

// recapitulando: aqui não há deadlock porque existem DUAS goroutines:
// a main e a anônima, dessa forma o envio e a recepção se encontram no canal, sincronizando as duas goroutines.
func main() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}
