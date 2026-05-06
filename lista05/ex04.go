// 4. Faça um programa que leia um vetor A de 10 posições, contendo números inteiros. Determine e mostre, a
// seguir, quais elementos de A estão repetidos e quantas vezes cada um se repete.

package main

import (
	f "fmt"
)

func main() {
	var a [10]int
	frequencia := make(map[int]int)

	// 1. Ler o vetor A de 10 posições
	f.Println("Digite 10 números inteiros:")
	for i := 0; i < 10; i++ {
		f.Printf("Posição %d: ", i+1)
		f.Scan(&a[i])
		// Conta a frequência de cada número automaticamente
		frequencia[a[i]]++
	}

	f.Println("\n--- Relatório de Repetições ---")

	// 2. Determinar e mostrar elementos repetidos e a contagem
	encontrouRepetido := false
	for numero, count := range frequencia {
		if count > 1 {
			f.Printf("O número %d se repete %d vezes.\n", numero, count)
			encontrouRepetido = true
		}
	}

	if !encontrouRepetido {
		f.Println("Não há elementos repetidos.")
	}
}
