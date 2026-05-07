// 19. Faça um programa que leia um primeiro vetor com dez números inteiros e um segundo vetor com cinco
// números inteiros. Mostre uma lista dos números do primeiro vetor com seus respectivos divisores armazenados
// no segundo vetor, bem como suas posições.
// Ex.: Num [ 5 12 4 7 10 3 2 6 23 16 ]
// Divis [ 3 11 5 8 2]
// Saída:
// Número 5:
// Divisível por 5 na posição 2
// Número 12:
// Divisível por 3 na posição 0
// Divisível por 2 na posição 4
// Número 4:
// Divisível por 2 na posição 4
// . . .

package main

import "fmt"

func main() {
	// Correção: Adicionado o tipo 'int' após o tamanho dos vetores
	var num [10]int
	var divis [5]int

	// 1. Leitura do primeiro vetor (10 números)
	fmt.Println("Digite 10 números inteiros:")
	for i := 0; i < 10; i++ {
		fmt.Printf("Vetor 1 [%d]: ", i)
		fmt.Scan(&num[i])
	}

	// 2. Leitura do segundo vetor (5 divisores)
	fmt.Println("\nDigite 5 números divisores:")
	for i := 0; i < 5; i++ {
		fmt.Printf("Vetor 2 [%d]: ", i)
		fmt.Scan(&divis[i])
	}

	fmt.Println("\n--- Resultado da Análise ---")

	// 3. Lógica de verificação de divisores
	for i := 0; i < 10; i++ {
		fmt.Printf("Número %d:\n", num[i])
		encontrou := false

		for j := 0; j < 5; j++ {
			// Verifica se o divisor não é zero para evitar erro de runtime
			if divis[j] != 0 && num[i]%divis[j] == 0 {
				fmt.Printf("  Divisível por %d na posição %d\n", divis[j], j)
				encontrou = true
			}
		}

		if !encontrou {
			fmt.Println("  Não possui divisores no segundo vetor.")
		}
	}
}
