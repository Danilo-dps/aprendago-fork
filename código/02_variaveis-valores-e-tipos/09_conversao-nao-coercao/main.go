package main

import "fmt"

type hotdog int

var b hotdog = 101

func main() {
	x := 10
	fmt.Printf("%v\n", x)
	fmt.Printf("%v\n", b)

	x = int(b) // isso funciona, apesar do b subjacente de hotdog ser int, o compilador não permite a conversão implícita de tipos, então é necessário fazer uma conversão explícita
	fmt.Printf("%v\n", x)
}
