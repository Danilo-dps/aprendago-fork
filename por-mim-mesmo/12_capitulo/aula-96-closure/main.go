package main

import "fmt"

func main() {
	a := i()

	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

	b := i()

	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
}

// closure: a função interna tem acesso ao escopo (variáveis) da função externa,
// mesmo depois que a função externa já retornou.
// isso acontece porque o Go detecta que 'x' é usada pela função interna e
// então aloca 'x' no heap (em vez da stack), fazendo com que ela sobreviva
// enquanto a closure que a referencia existir.
// cada chamada de i() cria uma nova variável 'x' independente, por isso
// 'a' e 'b' têm contadores separados, cada um mantendo seu próprio estado.
func i() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
