package main

import (
	"fmt"
	"strings"
)

func contarPalavras(s string) map[string]int {
	palavras := strings.Fields(s)
	ocorrencias := make(map[string]int)

	for _, palavra := range palavras {
		ocorrencias[palavra]++
	}

	return ocorrencias
}

func main() {
	texto := "o gato é um animal de estimação o gato é fofo"
	resultado := contarPalavras(texto)

	fmt.Println("Ocorrência de palavras:")
	for palavra, ocorrencia := range resultado {
		fmt.Printf("%s: %d\n", palavra, ocorrencia)
	}
}
