// 6. Escreva um programa que armazene todos os números inteiros de 100 a 1 (ordem decrescente) em um vetor. A
// seguir, imprima os elementos do vetor.

package main

import f "fmt"

func main() {
	const tamanho = 100

	numeros := make([]int, tamanho)

	// Preenche o vetor com números de 100 a 1
	for i := 0; i < tamanho; i++ {
		// Na primeira iteração (i=0), valor = 100 - 0 = 100
		// Na última iteração (i=99), valor = 100 - 99 = 1
		numeros[i] = 100 - i
	}

	f.Println("Vetor de 100 a 1:")
	for _, valor := range numeros {
		f.Printf("%d ", valor)
	}
	f.Println()
}
