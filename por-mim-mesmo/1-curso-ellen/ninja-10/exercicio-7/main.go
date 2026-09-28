package main

import "fmt"

// main cria um único canal bidirecional e o passa para duas funções com papéis
// diferentes: envia (produtor) e recebe (consumidor). envia roda em uma goroutine
// própria, então main segue direto para recebe.
func main() {
	canal := make(chan int)
	go envia(canal)
	recebe(canal)
}

// envia dispara 10 goroutines, cada uma enviando os valores de 0 a 9 no canal
// (100 envios no total). Ela mesma roda em uma goroutine (go envia), mas retorna
// logo após disparar as 10 internas, que continuam enviando por conta própria.
func envia(c chan<- int) {
	for range 10 {
		go func() {
			for j := range 10 {
				c <- j
			}
		}()
	}
}

// recebe faz exatamente 100 recebimentos (for range 100 é sobre um inteiro, não
// sobre o canal), o que bate com o total de envios (10 goroutines x 10 valores).
// Por isso não precisa de close: o loop termina pela contagem, não pelo
// fechamento. A ordem dos valores não é garantida, pois as 10 goroutines
// competem pelo mesmo canal.
func recebe(r <-chan int) {
	for a := range 100 {
		fmt.Println(a, "\t", <-r)
	}
}
