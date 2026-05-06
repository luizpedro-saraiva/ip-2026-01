// 8. Escreva um programa que receba quinze números inteiros e armazene em um vetor a raiz quadrada de cada
// número. Caso o valor digitado seja menor do que zero, o número -1 deve ser atribuído ao elemento do vetor.
// Após isso, imprima todos os valores armazenados.

package main

import (
	f "fmt"
	"math"
)

func main() {
	var numeros [15]int
	var resultados [15]float64

	f.Println("Digite 15 números inteiros:")

	for i := 0; i < 15; i++ {
		f.Printf("%dº número: ", i+1)
		f.Scan(&numeros[i])

		if numeros[i] < 0 {
			// Atribui -1 se o número for negativo
			resultados[i] = -1
		} else {
			// Calcula a raiz quadrada (convertendo int para float64)
			resultados[i] = math.Sqrt(float64(numeros[i]))
		}
	}

	f.Println("\nValores armazenados no vetor de resultados:")
	for i, valor := range resultados {
		if valor == -1 {
			f.Printf("Posição %d: %.0f (Número original negativo)\n", i, valor)
		} else {
			f.Printf("Posição %d: %.2f\n", i, valor)
		}
	}
}
