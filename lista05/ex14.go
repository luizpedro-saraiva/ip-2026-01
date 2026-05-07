// 14. Faça um programa que leia dois vetores de 10 elementos inteiros cada um e mostre o vetor resultante da
// intercalação desses dois vetores.
// Exemplo: Vetor 1 [0 5 4 2 1 5 3 2 5 9]
// Vetor 2 [1 5 4 2 0 5 3 2 5 9]
// Vetor resultante da intercalação [0 1 5 5 4 4 2 2 1 0 5 5 3 3 2 2 5 5 9 9]
package main

import "fmt"

func main() {

	vetor1 := [10]int{0, 5, 4, 2, 1, 5, 3, 2, 5, 9}
	vetor2 := [10]int{1, 5, 4, 2, 0, 5, 3, 2, 5, 9}

	var resultado [20]int

	for i := 0; i < 10; i++ {
		// O primeiro elemento vai para a posição par (2*i)
		resultado[2*i] = vetor1[i]
		// O segundo elemento vai para a posição ímpar (2*i + 1)
		resultado[2*i+1] = vetor2[i]
	}

	fmt.Println("Vetor 1:", vetor1)
	fmt.Println("Vetor 2:", vetor2)
	fmt.Println("Vetor resultante da intercalação:")
	fmt.Println(resultado)
}
