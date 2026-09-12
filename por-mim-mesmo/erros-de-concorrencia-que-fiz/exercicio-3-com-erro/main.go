package main

import (
	"fmt"
	"runtime"
	"sync"
)

// OBJETIVO DO EXERCÍCIO:
// Criar uma race condition de propósito, disparando várias goroutines que leem e escrevem
// nas variáveis globais `contador` e `contadorInterno` ao mesmo tempo, sem nenhuma sincronização
// (sem mutex, sem atomic), para observar na prática o comportamento imprevisível gerado por isso.
//
// O QUE EU TENTEI FAZER:
//   - Disparar 6 goroutines de incrementar() intercaladas com 6 goroutines de valorAtualContador()
//   - Usar runtime.Gosched() para forçar o escalonador a alternar entre goroutines e aumentar
//     a chance de ver a race condition acontecendo na prática
//   - Usar wg.Add(totalDeGoroutines) e wg.Wait() para esperar todas as 12 goroutines terminarem
//     antes do programa encerrar
//
// O QUE DEU ERRADO (E POR QUÊ):
//
//  1. Às vezes o programa termina sem imprimir todos os "Valor Atual"
//     Isso acontece porque só incrementar() chama wg.Done(). valorAtualContador() nunca chama
//     Done(), então o WaitGroup não sabe que precisa esperar por ela. Ou seja, wg.Add(totalDeGoroutines)
//     conta 12, mas só 6 Done() vão ser chamados (um por incrementar()). Isso já é um bug por si só,
//     mas ele fica mascarado quando totalDeGoroutines bate certo com o número de chamadas a incrementar().
//
//  2. Quando aumento totalDeGoroutines, o programa trava com "deadlock!"
//     Esse é o mesmo bug do item 1, só que agora ele fica visível: wg.Add(totalDeGoroutines) passa
//     a esperar um número de Done() maior do que o número de vezes que incrementar() é realmente
//     chamado. O contador interno do WaitGroup nunca chega a zero, então wg.Wait() bloqueia
//     a goroutine main() para sempre. O runtime detecta que nenhuma goroutine consegue mais
//     progredir e aborta com "all goroutines are asleep - deadlock!" — não é que as goroutines
//     "dormiram por acaso", é que elas estão todas bloqueadas esperando algo que nunca vai acontecer.
//
//  3. O valor final de `contador` muda a cada execução (race condition, intencional)
//     contador e contadorInterno são globais e compartilhadas por todas as goroutines de
//     incrementar(), que fazem `contadorInterno += i` e `contador = contadorInterno` sem
//     nenhum lock. Múltiplas goroutines leem e escrevem essas variáveis ao mesmo tempo,
//     então incrementos podem se sobrescrever ou se perder dependendo da ordem de execução
//     escolhida pelo escalonador — por isso o resultado é diferente a cada rodada.
//
// LIÇÃO:
// Esse código serve como exemplo do que NÃO fazer: WaitGroup mal configurado (Add não bate
// com o número real de Done()) e acesso concorrente a variáveis compartilhadas sem sincronização.
var wg sync.WaitGroup
var contador = 0
var contadorInterno = contador

func main() {

	totalDeGoroutines := 12
	wg.Add(totalDeGoroutines)

	go incrementar()
	runtime.Gosched()
	go valorAtualContador()
	runtime.Gosched()

	go incrementar()
	runtime.Gosched()
	go valorAtualContador()
	runtime.Gosched()

	go incrementar()
	runtime.Gosched()
	go valorAtualContador()
	runtime.Gosched()

	go incrementar()
	runtime.Gosched()
	go valorAtualContador()
	runtime.Gosched()

	go incrementar()
	runtime.Gosched()
	go valorAtualContador()
	runtime.Gosched()

	go incrementar()
	runtime.Gosched()
	go valorAtualContador()
	runtime.Gosched()
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
}
