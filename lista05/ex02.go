// Faça um programa que preencha um vetor com 10 números inteiros e um segundo vetor com 5 elementos
// inteiros, ambos informados pelo usuário. Calcule e mostre dois vetores resultantes. O primeiro vetor
// resultante
// será composto pelos números pares do primeiro vetor somado a todos os elementos do segundo vetor. O
// segundo será composto pelos números ímpares do primeiro vetor somado a todos os elementos do segundo
// vetor.
// Exemplo:
// Primeiro vetor: [ 4 7 5 8 2 15 9 6 10 11 ]
// Segundo vetor: [ 3 4 5 8 2 ]
// Primeiro vetor resultante: [ 26 30 24 . . . ], onde 26 = 4+3+4+5+8+2, 30 = 8+3+4+5+8+2, ...
// Segundo vetor resultante: [ 29 27 37 . . .], onde 29 =7+3+4+5+8+2, 27 = 5+3+4+5+8+2, ...

package main

import (
	f "fmt"
)

func main() {
	// Declaração dos vetores
	vetor1 := make([]int, 10)
	vetor2 := make([]int, 5)

	f.Println("Digite 10 números inteiros para o primeiro vetor:")
	for i := 0; i < 10; i++ {
		f.Scan(&vetor1[i])
	}

	f.Println("Digite 5 números inteiros para o segundo vetor:")
	for i := 0; i < 5; i++ {
		f.Scan(&vetor2[i])
	}

	// Calcula a soma de todos os elementos do segundo vetor
	somaVetor2 := 0
	for _, valor := range vetor2 {
		somaVetor2 += valor
	}

	// Criação dos vetores resultantes (usando slices, pois o tamanho pode variar se não houver pares/ímpares)
	var paresResultante []int
	var imparesResultante []int

	for _, valor := range vetor1 {
		if valor%2 == 0 {
			// É par: soma ao primeiro vetor resultante
			paresResultante = append(paresResultante, valor+somaVetor2)
		} else {
			// É ímpar: soma ao segundo vetor resultante
			imparesResultante = append(imparesResultante, valor+somaVetor2)
		}
	}

	// Exibição dos resultados
	f.Println("\n--- Resultados ---")
	f.Println("Vetor 1:", vetor1)
	f.Println("Vetor 2:", vetor2)
	f.Println("Soma do Vetor 2:", somaVetor2)
	f.Println("Primeiro vetor resultante (Pares + Sum V2):", paresResultante)
	f.Println("Segundo vetor resultante (Ímpares + Sum V2):", imparesResultante)
}
