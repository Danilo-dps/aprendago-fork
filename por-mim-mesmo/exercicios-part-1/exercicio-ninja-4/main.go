package main

import "fmt"

type ninjaOne int

var x ninjaOne

func main() {
	fmt.Printf("%v, %T", x, x)

	x = 42

	fmt.Printf("\n%v, %T\n", x, x)
}
