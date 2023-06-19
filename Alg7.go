package main

import (
	"fmt"
	"strings"
)

func contarLetrasPorPalavra(frase string) map[string]map[rune]int {
	contagem := make(map[string]map[rune]int)

	palavras := strings.Fields(frase)

	for _, palavra := range palavras {
		letras := make(map[rune]int)

		for _, letra := range palavra {
			letras[letra]++
		}

		contagem[palavra] = letras
	}

	return contagem
}

func main() {
	frase := "Aprender programação é divertido"

	resultado := contarLetrasPorPalavra(frase)

	fmt.Println("Contagem de letras por palavra:")
	for palavra, letras := range resultado {
		fmt.Printf("%s: %v\n", palavra, letras)
	}
}
