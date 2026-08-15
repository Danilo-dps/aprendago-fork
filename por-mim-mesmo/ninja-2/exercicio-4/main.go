package main

import "fmt"

func main() {
	num1 := 200

	fmt.Printf("%b\t%d\t%#x\n", num1, num1, num1)

	num2 := num1 << 1

	fmt.Printf("%b\t%d\t%#x\n", num2, num2, num2)

	num3 := num2 >> 1

	fmt.Printf("%b\t%d\t%#x\n", num3, num3, num3)
}
