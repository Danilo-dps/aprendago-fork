package main

import "fmt"

func main() {
	minhaStructAnonima := struct {
		mepe map[string]int
		ss   []string
	}{
		mepe: map[string]int{
			"golang": 2006,
			"java":   1995,
			"C":      1972,
		},
		ss: []string{"hello world", "gopher"},
	}

	fmt.Println(minhaStructAnonima)
}
