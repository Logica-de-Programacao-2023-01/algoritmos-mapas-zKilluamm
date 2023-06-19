package main

import "fmt"

func somarValores(m mapaInteiros) int {
	soma := 0

	for _, valor := range m {
		soma += valor
	}

	return soma
}

type mapaInteiros map[string]int

func main() {
	mapa := mapaInteiros{
		"chave1": 10,
		"chave2": 20,
		"chave3": 30,
	}

	resultado := somarValores(mapa)

	fmt.Println("Soma dos valores:", resultado)
}
