package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println()
	fmt.Println("Lista desordenada de ints")
	fmt.Println(xi)

	// sort xi
	fmt.Println()
	sort.Ints(xi)
	fmt.Println("Lista ordenada de ints")
	fmt.Println(xi)

	fmt.Println()
	fmt.Println("Lista desordenada de strings")
	fmt.Println(xs)

	// sort xs
	fmt.Println()
	fmt.Println("Lista ordenada de strings")
	sort.Strings(xs)
	fmt.Println(xs)
	fmt.Println()
}
