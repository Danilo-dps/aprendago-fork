package main

import "fmt"

func main() {
	anoNascimento := 1996
	anoAtual := 2026

	for anoNascimento <= anoAtual {
		fmt.Printf("%d\n", anoNascimento)
		anoNascimento++
	}
}
