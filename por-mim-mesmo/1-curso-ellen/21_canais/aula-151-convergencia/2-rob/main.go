package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	canal := converge(trabalho("maçã"), trabalho("pêra"))
	// for range 16 apenas repete o loop 16 vezes; não precisamos do índice,
	// só queremos consumir 16 mensagens do canal (a ordem entre "maçã" e "pêra"
	// não é definida, depende de qual goroutine enviar primeiro)
	for range 16 {
		fmt.Println(<-canal)
	}
}

func trabalho(s string) chan string {
	canal := make(chan string)

	go func(s string, c chan string) {
		// aqui o for é infinito, porque a intenção é que ele continue enviando dados para o canal
		// e nesse fluxo quem define o termino é o for range do main, que vai percorrer 16 vezes
		// é por isso que aqui não tem condição de parada
		for i := 1; ; i++ {
			c <- fmt.Sprintf("Função %v diz: %v", s, i)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
		}
	}(s, canal)
	return canal
}

// aqui `<-x` recebe um valor do canal x, e `novo <-` envia esse valor
// para o canal novo, tudo em um único statement (equivale a: v := <-x; novo <- v).
// Isso é o padrão fan-in: uma goroutine para cada canal de origem (x e y),
// cada uma repassando o que recebe para o canal novo compartilhado.
func converge(x, y chan string) chan string {
	novo := make(chan string)
	go func() {
		for {
			novo <- <-x
		}
	}()
	go func() {
		for {
			novo <- <-y
		}
	}()
	return novo
}
