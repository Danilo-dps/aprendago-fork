package main

import "fmt"

const (
	a int = 42
	b
	c bool = true
	d
	e string = "marmotinha gopher golang"
	f
)

func main() {
	s := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v", a, b, c, d, e, f)

	fmt.Println(s)
}
