// 17. Escreva um programa que leia um vetor de dez elementos inteiros e mostre os números primos e suas
// respectivas posições.

package main

import (
	"fmt"
)

// Função auxiliar para verificar se um número é primo
func ehPrimo(n int) bool {
	if n <= 1 {
		return false
	}
	// Verifica divisores de 2 até a raiz quadrada de n
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var vetor [10]int

	// 1. Lendo os dez elementos inteiros
	fmt.Println("Digite 10 números inteiros:")
	for i := 0; i < 10; i++ {
		fmt.Printf("Elemento %d: ", i+1)
		fmt.Scan(&vetor[i])
	}

	fmt.Println("\n--- Números Primos e suas Posições ---")

	// 2. Verificando e mostrando os números primos e posições
	encontrouPrimo := false
	for i := 0; i < 10; i++ {
		if ehPrimo(vetor[i]) {
			fmt.Printf("Primo: %d | Posição (índice): %d\n", vetor[i], i)
			encontrouPrimo = true
		}
	}

	if !encontrouPrimo {
		fmt.Println("Nenhum número primo foi encontrado no vetor.")
	}
}
