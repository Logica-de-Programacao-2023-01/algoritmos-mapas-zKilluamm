package main

import "fmt"

func contarCaracteres(s string) map[rune]int {
	frequencia := make(map[rune]int)

	for _, char := range s {
		frequencia[char]++
	}

	return frequencia
}

func main() {
	texto := "Hello, world!"

	resultado := contarCaracteres(texto)

	fmt.Println("Frequência de caracteres:")
	for char, frequencia := range resultado {
		fmt.Printf("%c: %d\n", char, frequencia)
	}
}
