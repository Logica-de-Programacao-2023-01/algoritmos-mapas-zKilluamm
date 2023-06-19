package main

import (
	"fmt"
	"sort"
	"strings"
)

type ByLetters []string

func (a ByLetters) Len() int           { return len(a) }
func (a ByLetters) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByLetters) Less(i, j int) bool { return sortString(a[i]) < sortString(a[j]) }

func sortString(s string) string {
	letters := strings.Split(s, "")
	sort.Strings(letters)
	return strings.Join(letters, "")
}

func encontrarAnagramas(palavras []string) map[string][]string {
	anagramas := make(map[string][]string)

	for _, palavra := range palavras {
		letrasOrdenadas := sortString(palavra)
		anagramas[letrasOrdenadas] = append(anagramas[letrasOrdenadas], palavra)
	}

	return anagramas
}

func main() {
	palavras := []string{"amor", "roma", "casa", "saco", "maçã", "çama"}

	resultado := encontrarAnagramas(palavras)

	fmt.Println("Grupos de palavras anagramas:")
	for _, grupo := range resultado {
		fmt.Println(grupo)
	}
}
