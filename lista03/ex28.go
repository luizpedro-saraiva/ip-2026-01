// 28. O valor aproximado de Π pode ser calculado usando-se a série:
// S = 1/1³ − 1/3³ + 1/5³ − 1/7³ + 1/9³ − ...

// sendo Π = ³√(S×32)

// Faça um programa que calcule e escreva o valor de Π com 51 termos.

package main

import (
	f "fmt"
	"math"
)

func main() {

	var soma float64 = 0.0
	sinal := 1.0

	num := 1.0

	for i := 0; i < 51; i++ {

		termo := 1.0 / (num * num * num)
		soma += sinal * termo

		// prepara próximo termo
		num += 2
		sinal *= -1
	}

	pi := math.Cbrt(32 * soma)

	f.Printf("Valor aproximado de PI: %.6f\n", pi)
}
