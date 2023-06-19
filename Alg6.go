package main

import "fmt"

func somarContagens(contagens []map[string]int) map[string]int {
	soma := make(map[string]int)

	for _, contagem := range contagens {
		for palavra, quantidade := range contagem {
			soma[palavra] += quantidade
		}
	}

	return soma
}

func main() {
	texto1 := map[string]int{
		"olá":      2,
		"mundo":    1,
		"gato":     3,
		"cachorro": 4,
	}

	texto2 := map[string]int{
		"olá":      1,
		"gato":     2,
		"cachorro": 2,
		"rato":     1,
	}

	texto3 := map[string]int{
		"gato":     1,
		"cachorro": 3,
		"rato":     2,
		"elefante": 1,
	}

	contagens := []map[string]int{texto1, texto2, texto3}

	resultado := somarContagens(contagens)

	fmt.Println("Soma das contagens:")
	for palavra, quantidade := range resultado {
		fmt.Printf("%s: %d\n", palavra, quantidade)
	}
}
