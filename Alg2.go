package main

import "fmt"

func mesclarMapas(mapa1, mapa2 map[string]int) map[string]int {
	resultado := make(map[string]int)

	// Copiar os elementos do mapa1 para o resultado
	for chave, valor := range mapa1 {
		resultado[chave] = valor
	}

	// Adicionar ou atualizar os elementos do mapa2 no resultado
	for chave, valor := range mapa2 {
		resultado[chave] = valor
	}

	return resultado
}

func main() {
	mapa1 := map[string]int{
		"chave1": 10,
		"chave2": 20,
	}

	mapa2 := map[string]int{
		"chave2": 30,
		"chave3": 40,
	}

	resultado := mesclarMapas(mapa1, mapa2)

	fmt.Println("Mapa Resultado:")
	for chave, valor := range resultado {
		fmt.Printf("%s: %d\n", chave, valor)
	}
}
