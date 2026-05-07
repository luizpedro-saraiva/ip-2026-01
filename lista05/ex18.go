// 18. Faça um programa que leia um vetor de elementos inteiros de 10 posições, mas já o leia de maneira ordenada
// crescente. Ao final, imprima o vetor.

package main

import "fmt"

func main() {
	var vetor [10]int

	fmt.Println("Digite 10 números inteiros em ordem crescente:")

	for i := 0; i < 10; i++ {
		var num int
		fmt.Printf("Elemento %d: ", i)
		fmt.Scan(&num)

		// Verifica se o número atual é menor que o anterior
		if i > 0 && num < vetor[i-1] {
			fmt.Printf("Inválido! Digite um valor maior ou igual a %d.\n", vetor[i-1])
			i-- // Decrementa para repetir a leitura desta posição
			continue
		}

		vetor[i] = num
	}

	fmt.Printf("\nVetor lido: %v\n", vetor)
}
