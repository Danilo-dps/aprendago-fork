package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Meu OS é:\t", runtime.GOOS)
	fmt.Println("Minha ARCH é:\t", runtime.GOARCH)
}
