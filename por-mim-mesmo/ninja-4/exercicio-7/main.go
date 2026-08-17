package main

import "fmt"

func main() {
	x := [][]string{
		{"joão", "pereira", "futebol"},
		{"maria", "pereira", "handbol"},
		{"lucas", "silva", "volei"},
	}

	for _, value := range x {
		fmt.Println(value[0])
		for _, item := range value {
			fmt.Println("\t", item)
		}
	}
}
