package main

import "fmt"

func main() {
	slice := []string{"banana", "maçã", "jaca", "pêssego"}

	//for índice, valor := range slice {fmt.Println("No índice", índice, "temos o valor:", valor)}

	//slice[4] = "melancia"
	slice = append(slice, "melancia")

	for _, valor := range slice {
		fmt.Printf("Um dos valores desse slice é %v.\n", valor)
	}
}
