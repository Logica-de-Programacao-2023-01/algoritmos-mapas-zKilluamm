package main

import "fmt"

func contarPares(slice []int) map[[2]int]int {
	frequencia := make(map[[2]int]int)

	// Conta a frequência dos pares de números
	for i := 0; i < len(slice)-1; i++ {
		for j := i + 1; j < len(slice); j++ {
			pair := [2]int{slice[i], slice[j]}
			frequencia[pair]++
		}
	}

	return frequencia
}

func main() {
	numeros := []int{1, 2, 3, 4, 5, 1, 2, 3}

	resultado := contarPares(numeros)

	fmt.Println("Frequência de pares:")
	for par, frequencia := range resultado {
		fmt.Printf("%v: %d\n", par, frequencia)
	}
}
