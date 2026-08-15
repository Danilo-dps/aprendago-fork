package main

import "fmt"

// em go é possível a função retornar mais de um valor,
// como no caso do Println que retorna o número de bytes escritos e um erro caso ocorra algum problema na escrita.
func main() {
	numerosdebytes, erros := fmt.Println("Hello, world!", "Oi galera", 100)
	fmt.Println(numerosdebytes, erros)
}
