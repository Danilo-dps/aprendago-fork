package main

import "fmt"

// func (receiver) identifier (parameters) (returns) { code }
func minhaFunc1(x ...int) int {
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma
}

func minhaFunc2(x []int) int {
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma
}

func main() {
	a := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25}
	f1 := minhaFunc1(a...)
	f2 := minhaFunc2(a)

	fmt.Println(f1)
	fmt.Println(f2)
}
