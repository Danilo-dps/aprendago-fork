package main

import "fmt"

var (
	a int  = 10
	b int  = 3
	c bool = a == b
	d bool = a != b
	e bool = a < b
	f bool = a > b
	g bool = a <= b
	h bool = a >= b
)

func main() {
	fmt.Printf("valor de a: %v\n", a)
	fmt.Printf("valor de b: %v\n", b)
	fmt.Printf("a é igual a b? %v\n", c)
	fmt.Printf("a é diferente de b? %v\n", d)
	fmt.Printf("a é menor que b? %v\n", e)
	fmt.Printf("a é maior que b? %v\n", f)
	fmt.Printf("a é menor ou igual a b? %v\n", g)
	fmt.Printf("a é maior ou igual a b? %v\n", h)
}
