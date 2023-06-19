package main

import (
	"fmt"
)

func calcularSaldo(contas map[string]float64) map[string]float64 {
	saldo := make(map[string]float64)
	total := 0.0

	// Calcula o total gasto
	for _, valor := range contas {
		total += valor
	}

	// Calcula o saldo de cada pessoa
	media := total / float64(len(contas))
	for pessoa, valor := range contas {
		saldo[pessoa] = valor - media
	}

	return saldo
}

func main() {
	contas := map[string]float64{
		"João":   100.0,
		"Maria":  50.0,
		"Pedro":  80.0,
		"Carlos": 120.0,
	}

	saldo := calcularSaldo(contas)

	fmt.Println("Saldo por pessoa:")
	for pessoa, valor := range saldo {
		fmt.Printf("%s: %.2f\n", pessoa, valor)
	}
}
