package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var contador = 0
var contadorInterno = contador

// Esse código está com race condition
func main() {

	totalDeGoroutines := 12
	wg.Add(totalDeGoroutines)

	for range totalDeGoroutines / 2 {
		go incrementar()
		runtime.Gosched()
		go valorAtualContador()
		runtime.Gosched()
	}
	wg.Wait()
}

func incrementar() {
	for i := range 100 {
		contadorInterno += i
	}

	contador = contadorInterno
	wg.Done()
}

func valorAtualContador() {
	fmt.Println("Valor Atual: ", contador)
	wg.Done()
}
