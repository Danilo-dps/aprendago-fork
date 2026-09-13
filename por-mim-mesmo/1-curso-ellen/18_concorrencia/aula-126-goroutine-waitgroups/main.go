package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// WaitGroup é usado para esperar um grupo de goroutines terminarem antes que o
// programa continue (ou finalize). Ele mantém um contador interno: cada goroutine
// que deve ser aguardada soma 1 nesse contador, e ao terminar chama Done() para
// decrementar. wg.Wait() bloqueia até o contador chegar a zero
var wg sync.WaitGroup

// goroutine é a forma do Go de executar código concorrentemente. Apesar do nome
// lembrar "thread", goroutines NÃO são threads do sistema operacional — são
// gerenciadas pelo runtime do Go, que as multiplexa sobre um número menor de
// threads reais do SO. Por isso são muito mais leves (memória inicial de poucos KB,
// contra MBs de uma thread do SO), permitindo criar milhares/milhões delas
func main() {

	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())

	// informa ao WaitGroup que deve aguardar 2 goroutines (incrementa o contador em 2)
	wg.Add(2)

	// a palavra-chave 'go' inicia uma nova goroutine para executar a função
	// concorrentemente, sem esperar ela terminar para seguir o fluxo
	go func1()
	go func2()

	fmt.Println(runtime.NumGoroutine())

	// bloqueia a goroutine main até o contador do WaitGroup chegar a zero,
	// ou seja, até func1 e func2 chamarem wg.Done()
	wg.Wait()

}

func func1() {
	for i := range 100 {
		fmt.Println("func1:", i)
		time.Sleep(20)
	}
	// sinaliza que a execução aqui finalizou
	wg.Done()
}

func func2() {
	for i := range 100 {
		fmt.Println("func2:", i)
		time.Sleep(20)
	}
	// sinaliza que a execução aqui finalizou
	wg.Done()
}

// concorrência, paralelismo e multithread são conceitos diferentes que muitas vezes
// aparecem juntos no mesmo processo ou cenário

// concorrência é quando o programa é estruturado para lidar com várias tarefas
// "ao mesmo tempo" (não necessariamente no mesmo instante), intercalando a execução
// entre elas. É a coisa mais comum na programação: várias tarefas competindo por
// uma fatia de tempo de processamento

// paralelismo é quando duas ou mais tarefas de fato são executadas no mesmo instante,
// o que só é possível havendo múltiplos núcleos de CPU disponíveis

// multithread é a capacidade do ambiente/processo de ter várias threads em execução.
// Um pool de threads é apenas uma técnica para gerenciar essas threads (reaproveitá-las
// em vez de criar e destruir a cada tarefa), não é sinônimo de multithread

// exemplo com os três conceitos no mesmo fluxo: meu servidor tem um pool de 1000 threads,
// e a aplicação recebe 1000 requisições por segundo. Supondo que a máquina tenha, por
// exemplo, 8 núcleos de CPU, no máximo 8 requisições conseguem ser processadas em
// paralelo em um dado instante. As demais threads ficam concorrendo entre si pelo tempo
// de CPU disponível. Ou seja, o cenário tem multithread (o pool), paralelismo (as poucas
// requisições rodando literalmente ao mesmo tempo) e concorrência (todas as outras
// disputando os núcleos disponíveis)
