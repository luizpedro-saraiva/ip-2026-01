// 15. Crie um programa que leia um vetor de 30 números inteiros e gere um segundo vetor cujas posições pares
// possuirão elementos que serão o dobro do vetor original e as ímpares, o triplo.
package main

import "fmt"

func main() {

	var original [30]int
	var resultado [30]int

	fmt.Println("Digite 30 números inteiros:")

	for i := 0; i < 30; i++ {
		fmt.Scan(&original[i])
	}

	// 2. Processamento conforme a posição (i)
	for i := 0; i < 30; i++ {
		if i%2 == 0 {

			resultado[i] = original[i] * 2
		} else {

			resultado[i] = original[i] * 3
		}
	}

	fmt.Println("\nVetor Original:", original)
	fmt.Println("Vetor Resultante:", resultado)
}
