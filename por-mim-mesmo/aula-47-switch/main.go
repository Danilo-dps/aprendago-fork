package main

import "fmt"

var x interface{}

func main() {
	x = true

	switch x.(type) {
	case int:
		fmt.Println("x é do tipo inteiro")
	case bool:
		fmt.Println("x é do tipo bool")
	case string:
		fmt.Println("x é do tipo string")
	case float64:
		fmt.Println("x é do tipo float64")
	default:
		fmt.Println("x é de outro tipo")
	}

	fmt.Println("===========================")

	switch x := 2; x {
	case 1:
		fmt.Println("x é do tipo 1")
	case 2:
		fmt.Println("x é do tipo 2")
	case 3:
		fmt.Println("x é do tipo 3")
	case 4:
		fmt.Println("x é do tipo 4")
	default:
		fmt.Println("x não é nenhum dos tipos acima")
	}
	
}
