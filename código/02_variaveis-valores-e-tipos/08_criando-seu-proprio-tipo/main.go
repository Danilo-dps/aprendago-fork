package main

import "fmt"

type hotdog int

var b hotdog = 10

func main() {
	x := 10
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", b)

	// b = x isso não funciona, apesar do b subjacente de hotdog ser int, o compilador não permite a conversão implícita de tipos, então é necessário fazer uma conversão explícita
	b = hotdog(x)
	fmt.Printf("%T\n", b)
}
