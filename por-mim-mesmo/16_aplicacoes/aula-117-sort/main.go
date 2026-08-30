package main

import (
	"fmt"
	"sort"
)

func main() {
	ss := []string{"abóbora", "maçã", "laranja", "beringela", "berinjela"}

	fmt.Println("Lista de String desordenada")
	fmt.Println(ss)

	sort.Strings(ss)

	fmt.Println("Lista de String ordenada")
	fmt.Println(ss)

	fmt.Println("======================================")

	si := []int{123, 987, 324, 876, 234, 987, 234, 76}

	fmt.Println("Lista de ints desordenada")
	fmt.Println(si)

	sort.Ints(si)

	fmt.Println("Lista de ints ordenada")
	fmt.Println(si)
}
