package main

import (
	"fmt"
)

var (
	a int
	b float64
	c string
	d bool
)

func main() {
	// Valores padrão para quando o dado é declarado padrão de quando o dado é não é inicialido
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", d, d)
}
