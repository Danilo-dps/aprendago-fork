package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

// correção de race condition usando atomic
var wg sync.WaitGroup
var contador int64
var contadorInterno = contador

func main() {
	atomic.SwapInt64(&contador, 0)
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

	defer wg.Done()

	for i := range 100 {
		atomic.AddInt64(&contadorInterno, int64(i))
	}

	atomic.SwapInt64(&contador, atomic.LoadInt64(&contadorInterno))
}

func valorAtualContador() {
	defer wg.Done()

	fmt.Println("Valor Atual: ", atomic.LoadInt64(&contador))
}
