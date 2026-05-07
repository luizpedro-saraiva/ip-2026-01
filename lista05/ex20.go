// 20. Faça um programa que receba o números sorteado em um dado durante 20 jogadas, mostre os números
// sorteados e a frequência com que apareceram.

package main

import "fmt"

func main() {
	var jogadas [20]int
	var frequencia [7]int // Índice 1 a 6 para facilitar a contagem

	fmt.Println("Digite o resultado de 20 jogadas (números de 1 a 6):")

	// 1. Recebe as jogadas e calcula a frequência
	for i := 0; i < 20; i++ {
		fmt.Printf("Jogada %d: ", i+1)
		fmt.Scan(&jogadas[i])

		// Valida se o número é uma face válida do dado
		if jogadas[i] < 1 || jogadas[i] > 6 {
			fmt.Println("Valor inválido! Digite entre 1 e 6.")
			i--
			continue
		}

		// Incrementa a frequência do número sorteado
		frequencia[jogadas[i]]++
	}

	// 2. Mostra os números sorteados
	fmt.Println("\nNúmeros sorteados:", jogadas)

	// 3. Mostra a frequência de cada face
	fmt.Println("\nFrequência de cada número:")
	for face := 1; face <= 6; face++ {
		fmt.Printf("Número %d: apareceu %d vez(es)\n", face, frequencia[face])
	}
}
