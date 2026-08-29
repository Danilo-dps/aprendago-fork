package main

import (
	"encoding/json"
	"fmt"
)

// tags em go: mapeiam o campo da struct para a chave do JSON.
// são essenciais quando os nomes divergem (ex: json em minúsculo,
// campo da struct em maiúsculo), mas aqui, como os nomes já
// coincidem, o Unmarshal funcionaria mesmo sem elas.
type informacoes struct {
	Nome          string  `json:"Nome"` // isso aqui são tags
	Sobrenome     string  `json:"Sobrenome"`
	Idade         int     `json:"Idade"`
	Profissao     string  `json:"Profissao"`
	Contabancaria float64 `json:"Contabancaria"`
}

// Marshal/Unmarshal trabalham com []byte em memória: você passa (ou recebe)
// os dados já prontos como slice de bytes, por isso existe essa variável
// intermediária (ex: 'sb' antes do Unmarshal).
// Encoder/Decoder trabalham direto com um io.Writer/io.Reader (ex: um arquivo
// ou uma conexão de rede), streamando os dados sem precisar montar tudo em
// memória antes — por isso não há essa etapa intermediária de []byte.
func main() {
	sb := []byte(`{"Nome":"James","Sobrenome":"Bond","Idade":40,"Profissao":"Agente Secreto","Contabancaria":1000000.5}`)

	var jamesbond informacoes
	err := json.Unmarshal(sb, &jamesbond)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(jamesbond)
	fmt.Println(jamesbond.Profissao)
}
