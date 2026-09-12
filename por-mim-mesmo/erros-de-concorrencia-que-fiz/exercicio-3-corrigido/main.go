package main

import (
	"fmt"
	"runtime"
	"sync"
)

// VERSÃO CORRIGIDA
//
// Essa é a versão corrigida do exemplo que está em exercicio-3-com-erro/main.go.
// Os dois problemas identificados foram:
//
//  1. DEADLOCK: o WaitGroup esperava mais chamadas de Done() do que realmente aconteciam.
//     wg.Add(totalDeGoroutines) contava as 12 goroutines (6 incrementar + 6 valorAtualContador),
//     mas só incrementar() chamava wg.Done(). Correção: agora TODAS as goroutines que o
//     WaitGroup precisa esperar chamam wg.Done() (usei defer wg.Done() em ambas as funções),
//     então a contagem do Add() sempre bate com a contagem real de Done().
//
//  2. RACE CONDITION: contador e contadorInterno eram lidas e escritas por múltiplas
//     goroutines ao mesmo tempo sem nenhuma sincronização. Correção: um sync.Mutex protege
//     o acesso a essas variáveis compartilhadas. Toda leitura e escrita agora acontece
//     dentro de mu.Lock() / mu.Unlock(), garantindo que só uma goroutine por vez acesse
//     os dados, eliminando a condição de corrida.
var (
	wg              sync.WaitGroup
	mu              sync.Mutex
	contador        = 0
	contadorInterno = contador
)

func main() {

	totalDeGoroutines := 6
	// Add(totalDeGoroutines * 2) porque agora tanto incrementar() quanto
	// valorAtualContador() chamam wg.Done() — a contagem precisa bater com
	// o número real de goroutines disparadas (6 + 6 = 12).
	wg.Add(totalDeGoroutines * 2)

	for range totalDeGoroutines {
		go incrementar()
		runtime.Gosched()
		go valorAtualContador()
		runtime.Gosched()
	}

	wg.Wait()

	// Como só a main() executa aqui (todas as goroutines já terminaram),
	// não precisa de lock para ler o valor final.
	fmt.Println("Valor final do contador:", contador)
}

func incrementar() {
	// defer garante que Done() sempre será chamado, mesmo se a função
	// tivesse um return antecipado ou um panic no meio do caminho.
	defer wg.Done()

	// Lock protege contadorInterno e contador de acesso concorrente.
	// Enquanto essa goroutine estiver com o lock, nenhuma outra
	// consegue ler ou escrever nessas variáveis, evitando a race condition.
	mu.Lock()
	defer mu.Unlock()

	for i := range 100 {
		contadorInterno += i
	}

	contador = contadorInterno
}

func valorAtualContador() {
	// Antes essa função não chamava wg.Done(), o que causava o deadlock
	// quando o número de goroutines aumentava. Agora ela também é
	// rastreada pelo WaitGroup.
	defer wg.Done()

	// Lock também é necessário aqui: essa função LÊ contador enquanto
	// incrementar() pode estar ESCREVENDO nele ao mesmo tempo. Sem o lock,
	// isso ainda seria uma race condition (leitura concorrente com escrita).
	mu.Lock()
	valor := contador
	mu.Unlock()

	fmt.Println("Valor Atual: ", valor)
}
