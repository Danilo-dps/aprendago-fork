package main

import "fmt"

type ninjaOne int

var x ninjaOne

var y int

func main() {
	y = int(x)

	fmt.Printf("%v, %T\n", y, y)
}
