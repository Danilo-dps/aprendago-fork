package main

import (
	"fmt"
	"runtime"
	"sync"
)

// correção de race condition usando Mutex
var wg sync.WaitGroup
var mu sync.Mutex
var contador = 0
var contadorInterno = contador

func main() {

	totalDeGoroutines := 12
	wg.Add(totalDeGoroutines)
	defer wg.Wait()

	for range totalDeGoroutines / 2 {
		go incrementar()
		runtime.Gosched()
		go valorAtualContador()
		runtime.Gosched()
	}

}

func incrementar() {

	mu.Lock()
	defer wg.Done()
	defer mu.Unlock()

	for i := range 100 {
		contadorInterno += i
	}

	contador = contadorInterno
}

func valorAtualContador() {
	mu.Lock()
	defer wg.Done()
	defer mu.Unlock()

	fmt.Println("Valor Atual: ", contador)
}
