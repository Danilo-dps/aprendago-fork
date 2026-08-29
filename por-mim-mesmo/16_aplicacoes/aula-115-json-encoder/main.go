package main

import (
	"encoding/json"
	"os"
)

type pessoa struct {
	Nome          string
	Sobrenome     string
	Idade         int
	Profissao     string
	Contabancaria float64
}

// json.NewEncoder recebe um io.Writer (aqui, os.Stdout) e escreve o JSON
// diretamente nele, sem precisar gerar um []byte intermediário na memória
// como acontece com o Marshal.
// Encode() já serializa a struct e escreve o resultado direto no destino,
// adicionando uma quebra de linha (\n) ao final.
// útil para streaming: escrever direto em arquivos, respostas HTTP, conexões
// de rede, etc., sem precisar montar o JSON inteiro antes de enviar.
func main() {
	jamesbond := pessoa{
		Nome:          "James",
		Sobrenome:     "Bond",
		Idade:         40,
		Profissao:     "Agente Secreto",
		Contabancaria: 1000000.50,
	}

	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(jamesbond)
}
