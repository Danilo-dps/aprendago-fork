package main

import "fmt"

func main() {
	estadosBrasileiros := make([]string, 27)
	estadosBrasileiros[0], estadosBrasileiros[1], estadosBrasileiros[2],
		estadosBrasileiros[3], estadosBrasileiros[4], estadosBrasileiros[5],
		estadosBrasileiros[6], estadosBrasileiros[7], estadosBrasileiros[8],
		estadosBrasileiros[9], estadosBrasileiros[10], estadosBrasileiros[11],
		estadosBrasileiros[12], estadosBrasileiros[13], estadosBrasileiros[14],
		estadosBrasileiros[15], estadosBrasileiros[16], estadosBrasileiros[17],
		estadosBrasileiros[18], estadosBrasileiros[19], estadosBrasileiros[20],
		estadosBrasileiros[21], estadosBrasileiros[22], estadosBrasileiros[23],
		estadosBrasileiros[24], estadosBrasileiros[25], estadosBrasileiros[26] =
		"Acre", "Alagoas", "Amapá", "Amazonas", "Bahia", "Ceará", "Espírito Santo",
		"Goiás", "Maranhão", "Mato Grosso", "Mato Grosso do Sul", "Minas Gerais", "Pará",
		"Paraíba", "Paraná", "Pernambuco", "Piauí", "Rio de Janeiro", "Rio Grande do Norte",
		"Rio Grande do Sul", "Rondônia", "Roraima", "Santa Catarina", "São Paulo", "Sergipe",
		"Tocantins", "Goiana"

	fmt.Println(estadosBrasileiros)
	fmt.Printf("Length: %d, Capacity: %d\n", len(estadosBrasileiros), cap(estadosBrasileiros))
}
