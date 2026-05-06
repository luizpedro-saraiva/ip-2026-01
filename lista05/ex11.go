// 11. Faça um programa que receba 100 valores numéricos, armazene-os em um vetor, calcule e imprima o valor do
// somatório dado a seguir:
// S=(b0-b99)elevado3 + (b1-b98)elevado3 + (b2-b97)elevado3 + . . . + (b49-b50)elevado3

package main

import (
	f "fmt"
)

func main() {

	var b [20]float64
	var soma float64

	f.Println("Digite os 20 valores:")

	// 2. Receber os 20 valores
	for i := 0; i < 20; i++ {
		f.Printf("Valor %d: ", i+1)
		f.Scan(&b[i])
	}

	// 3. Calcular o somatório
	// O loop vai até 9 (metade de 20)
	// i cresce (0, 1...9) e j diminui (19, 18...10)
	for i := 0; i < 10; i++ {
		j := 19 - i
		diferenca := b[i] - b[j]

		// Elevando ao cubo: (b[i] - b[j])³
		cubo := diferenca * diferenca * diferenca
		soma += cubo
	}

	f.Printf("\nO valor do somatório S para 20 valores é: %.2f\n", soma)
}
