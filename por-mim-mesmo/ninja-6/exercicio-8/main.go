package main

import "fmt"

// func (receiver) identifier (parameters) (returns) { code }
func ola() func(sout string) string {
	return func(sin string) string {
		return sin + " é uma string"
	}
}

func main() {
	f := ola()

	v := f("valor estanho")

	fmt.Println(v)
}
