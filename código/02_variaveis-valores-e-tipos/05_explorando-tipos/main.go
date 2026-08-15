package main

import "fmt"

var x int = 10

// var y int
// y = 10, isso da erro de sintaxe, porque para declaração usando var é preciso atribui valor ou na declaração ou dentro de code block
func main() {
	// x = 20.5 gera erro porque x já foi declarado como sendo int
	x = 20
	fmt.Println(x)
}
