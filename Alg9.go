package main

import "fmt"

func fibonacciMap(n int) map[int]int {
	fibMap := make(map[int]int)

	// Verifica se o valor de n é válido
	if n < 0 {
		return fibMap
	}

	// Gera a sequência de Fibonacci
	a, b := 0, 1
	for i := 0; i <= n; i++ {
		fibMap[i] = a
		a, b = b, a+b
	}

	return fibMap
}

func main() {
	n := 10

	fibonacci := fibonacciMap(n)

	fmt.Println("Sequência de Fibonacci:")
	for i, num := range fibonacci {
		fmt.Printf("Índice %d: %d\n", i, num)
	}
}
