package main

import "fmt"

func main() {
	// String literal interpretada, isso faz com que o \n seja interpretado como uma quebra de linha, ou seja, como uma ação
	x := "Oi bom dia\ncomo vai?\tespero que esteja tudo bem"

	// String literal crua, isso faz com que o \n seja interpretado como dois caracteres, uma barra invertida e um n
	y := `"Oi bom dia\ncomo vai?\tespero que esteja tudo bem"`

	fmt.Println(x)
	fmt.Println(y)

	a := "oi"
	b := "bom dia"

	// usando a função Sprint do pacote fmt para concatenar as strings,
	// essa função não exibe o valoor na tela, ela apenas retorna o valor concatenado, que é armazenado na variável z
	z := fmt.Sprint(a, " ", b)

	fmt.Println(z)
}
