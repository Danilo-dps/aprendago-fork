package main

import "fmt"

const (
	_ = 2026 + iota
	year2027
	year2028
	year2029
	year2030
	year2031
)

func main() {
	fmt.Printf("%v\t %v\t %v\t %v\t %v\n", year2027, year2028, year2029, year2030, year2031)
}
