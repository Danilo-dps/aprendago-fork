package main

import "fmt"

type veiculo struct {
	portas int
	cor    string
}

type caminhote struct {
	veiculo
	tracaoNasQuatro bool
}

type sedan struct {
	veiculo
	modeloLuxo bool
}

func main() {
	minhaCaminhote := caminhote{
		veiculo{
			portas: 4,
			cor:    "prata",
		},
		true,
	}

	meuSedan := sedan{
		veiculo{
			portas: 2,
			cor:    "azul",
		},
		true,
	}

	fmt.Println(minhaCaminhote)
	fmt.Println(meuSedan)

	fmt.Printf("%v\n", minhaCaminhote.cor)
	fmt.Printf("%v\n", meuSedan.cor)
}
