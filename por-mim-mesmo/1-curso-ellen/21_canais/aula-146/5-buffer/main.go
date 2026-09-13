package main

import (
	"fmt"
)

// Este exemplo é o par do anterior: mesmo código, buffer maior.
//
// make(chan int, 2) cria espaço para 2 valores.
//   - canal <- 42: ocupa 1/2, não bloqueia.
//   - canal <- 43: ocupa 2/2, ainda não bloqueia (cabe).
//
// Os dois envios cabem no buffer, então a goroutine main não trava.
// Em seguida, as duas recepções retiram os valores em ordem FIFO:
// 42 primeiro, 43 depois.
//
// Comparando com o exemplo anterior (cap=1, mesmo código): lá o
// segundo envio não cabia e dava deadlock. Aqui cabe → funciona.
//
// Regra: canal de capacidade N aceita até N envios sem recepção.

func main() {
	canal := make(chan int, 2)
	canal <- 42
	canal <- 43
	fmt.Println(<-canal)
	fmt.Println(<-canal)
}
