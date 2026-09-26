package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// esse é o padrão fan-out/fan-in: manda() alimenta canal1 sequencialmente,
// e outra() distribui ("fan-out") cada valor recebido para uma goroutine própria,
// que roda trabalho() com um delay aleatório e envia o resultado para canal2 ("fan-in").
// como as goroutines rodam concorrentemente e dormem tempos diferentes,
// a ordem de chegada em canal2 não é garantida ser a mesma ordem de envio em canal1.
func main() {
	canal1 := make(chan int)
	canal2 := make(chan int)

	go manda(20, canal1)
	go outra(canal1, canal2)

	for v := range canal2 {
		fmt.Println(v)
	}
}

// aqui é preenchido o canal1 com n valores (0 a n-1), e depois é fechado,
// o que permite que o `for range canal1` em outra() termine quando o canal esvaziar
func manda(n int, canal chan int) {
	for i := range n {
		canal <- i
	}
	close(canal)
}

func outra(canal1, canal2 chan int) {
	var wg sync.WaitGroup

	for v := range canal1 {
		wg.Add(1)
		go func(x int) {
			canal2 <- trabalho(x)
			wg.Done()
		}(v)
	}
	wg.Wait()
	close(canal2)
}

func trabalho(n int) int {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
	return n
}
