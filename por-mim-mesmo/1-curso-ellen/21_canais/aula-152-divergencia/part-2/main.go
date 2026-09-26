package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	canal1 := make(chan int)
	canal2 := make(chan int)
	funcoes := 5
	go manda(100, canal1)
	go outra(funcoes, canal1, canal2)
	for v := range canal2 {
		fmt.Println(v)
	}
}

func manda(n int, canal chan int) {
	for i := range n {
		canal <- i
	}
	close(canal)
}

// aqui n (definido por funcoes) goroutines concorrentes leem do MESMO canal1 e escrevem em canal2;
// como múltiplas goroutines competem para receber de canal1, cada valor é processado por
// apenas uma delas (nunca duplicado), mas a ordem de chegada em canal2 não é garantida
// ser a mesma ordem de envio em canal1.
//
// for range funcoes só dispara as 5 goroutines (uma vez cada) — quem define o ritmo de
// chegada dos resultados é o sleep fixo de 1s em trabalho(). Como o sleep não é mais
// aleatório (diferente do exemplo anterior), as 5 goroutines tendem a começar e terminar
// quase juntas, dando a impressão de que os resultados chegam "em levas de 5 em 5".
//
// wg.Go(f) (API nova do WaitGroup) já faz Add(1) + go f() + Done() automaticamente,
// substituindo o padrão manual wg.Add(1)/go func(){...; wg.Done()}() usado antes.
//
// é um exemplo de worker pool: um volume grande de dados processado por um número
// LIMITADO de goroutines (funcoes), ajustável conforme a necessidade de paralelismo.
func outra(funcoes int, canal1, canal2 chan int) {
	var wg sync.WaitGroup
	for range funcoes {
		wg.Go(func() {
			for v := range canal1 {
				canal2 <- trabalho(v)
			}
		})
	}
	wg.Wait()
	close(canal2)
}

func trabalho(n int) int {
	time.Sleep(time.Millisecond * 1000) // time.Duration(rand.Intn(1e3)))
	return n
}
