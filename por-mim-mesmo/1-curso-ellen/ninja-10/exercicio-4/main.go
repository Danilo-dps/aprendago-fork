package main

import "fmt"

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

// gen cria um canal bidirecional c e dispara uma goroutine que envia os valores
// de 0 a 99 nele e depois o fecha (a goroutine usa a variável local c, por isso
// pode enviar e chamar close). Em seguida, envia um 0 em q, sinalizando o fim
// do fluxo (o valor não importa, o que importa é o envio).
//
// q é recebido como chan<- int (só envia): quem chama passa um canal
// bidirecional, convertido implicitamente. Como o envio em q bloqueia até
// alguém receber, é preciso ter um receptor lendo q, senão a goroutine fica
// presa para sempre.
//
// O retorno é <-chan int: no return c, o canal bidirecional é convertido
// implicitamente para só-recebimento, então quem chama só consegue ler,
// não consegue enviar nem fechar o canal.
func gen(q chan<- int) <-chan int {
	c := make(chan int)
	go func() {
		for i := range 100 {
			c <- i
		}
		close(c)
		q <- 0
	}()
	return c
}

// receive consome c e q com select. Depois do close(c), receber de c retorna
// o valor zero imediatamente e para sempre, então esse case ficaria sempre pronto
// e imprimiria zeros espúrios até o select sortear o case de q.
// Por isso usamos v, ok := <-c: quando ok é false, atribuímos nil a c, o que
// desativa esse case (canal nil nunca fica pronto) e deixa apenas q,
// que sinaliza o fim do fluxo.
func receive(c <-chan int, q <-chan int) {
	for {
		select {
		case v, ok := <-c:
			if !ok {
				c = nil // canal fechado: um case com canal nil nunca fica pronto
				continue
			}
			fmt.Println(v)
		case <-q:
			return
		}
	}
}
