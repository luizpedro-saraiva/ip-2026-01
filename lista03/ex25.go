// Faça um programa que calcule e escreva o valor de N no seguinte somatório:
// S = 1/225 − 2/196 + 4/169 − 8/144 + ... + 16384/1

package main

import "fmt"

func main() {

	var soma float64 = 0.0
	var termo float64
	var potencia float64 = 1.0 // 2^0
	var sinal float64 = 1.0

	for k := 1; k <= 15; k++ {

		// denominador: 15², 14², ..., 1²
		den := float64((16 - k) * (16 - k))

		termo = potencia / den

		soma += sinal * termo

		// prepara próximo termo
		potencia *= 2 // dobra (2^k)
		sinal *= -1   // alterna sinal
	}

	fmt.Printf("Resultado do somatório S = %.6f\n", soma)
}
