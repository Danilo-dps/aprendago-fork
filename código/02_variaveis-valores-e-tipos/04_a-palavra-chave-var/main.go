package main

import "fmt"

var y = 10 // var é usado a nível de pacote porque o operador curto := não funciona fora de um code block

// var z int
// z = 10, isso da erro de sintaxe, porque para declaração usando var é preciso atribui valor ou na declaração ou dentro de code block
func main() {
	x := 20 // aqui é usado o operador curto porque está dentro de um code block, preferencialmente, usar o :=
	qualquercoisa(x)
}

func qualquercoisa(x int) {
	fmt.Println(y) // acessando a variavel package-level scope
	fmt.Println(x) // utilizando a variavel recebida como parâmetro da função
}
