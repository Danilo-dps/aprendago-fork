package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Númeoro de Threads físicas do meu Pc", runtime.NumCPU())
	wg.Add(2)

	// goroutine
	go primeiroHello()
	runtime.Gosched() // permite que a thread main se desprenda para acionar outra função

	go segundoHello()
	runtime.Gosched()

	wg.Wait()
}

func primeiroHello() {
	fmt.Println("Hello gopher goroutine")
	fmt.Println("Qual Goroutine aqui", runtime.NumGoroutine())
	time.Sleep(3 * time.Second) // sleep intencional para tentar forçar o programa a usar outra thread
	wg.Done()
}

func segundoHello() {
	fmt.Println("Hello gopher goroutine")
	fmt.Println("Qual Goroutine aqui", runtime.NumGoroutine())
	time.Sleep(3 * time.Second)
	wg.Done()
}
