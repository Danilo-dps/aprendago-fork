package main

import "fmt"

// func (receiver) identifier (parameters) (returns) { code }
func multiplica(x ...int) int {
	n := 1
	for _, v := range x {
		n *= v
	}
	return n
}

func somentePares(f func(x ...int) int, y ...int) int {
	var slice []int
	for _, v := range y {
		if v%2 == 0 {
			slice = append(slice, v)
		}
	}
	total := f(slice...)
	return total
}

func main() {
	numerosPares := somentePares(multiplica, []int{50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60}...)
	fmt.Println(numerosPares)
}
