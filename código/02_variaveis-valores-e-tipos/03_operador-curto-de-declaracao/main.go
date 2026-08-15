package main

import "fmt"

var a = "Escopo de package"

func main() {
	/*
		:= parece uma marmota (gopher) ou o punisher.
		Uso:
		Tipagem automática
		Só pode repetir se houverem variáveis novas
		!= do assignment operator (operador de atribuição)
		Só funciona dentro de codeblocks
		Terminologia:
		keywords (palavras-chave) são termos reservados
		operadores, operandos
		statement (declaração, afirmação) → uma linha de código, uma instrução que forma uma ação, formada de expressões
		expressão → qualquer coisa que "produz um resultado"
		scope (abrangência)
		package-level scope
		Lição principal:
		:= utilizado pra criar novas variáveis, dentro de code blocks
		= para atribuir valores a variáveis já existentes */

	x := 10            // aqui é declarado e inicializado
	y := "Olá bom dia" // aqui também é declarado e inicializado

	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n", y, y)

	x, z := 20, 30 // aqui z é declarado e inicializado, enquanto x é reatribuído

	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n", y, y)
	fmt.Printf("z: %v, %T\n", z, z)
	fmt.Printf("a: %v, %T\n", a, a)
}
