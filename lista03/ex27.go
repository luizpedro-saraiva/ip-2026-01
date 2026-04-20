// 27. Faça um programa que:
// a) calcule o valor do cosseno de x através dos 20 elementos da série seguinte:
// coseno(x) = 1 − (x²/2!) + (x⁴/4!) − (x⁶/6!) + (x⁸/8!) − ...

// b) calcule a diferença entre o valor calculado em a e o valor fornecido por função
// c) imprima o que foi calculado nos itens a e b.

// Observação: o valor de x é fornecido como entrada.

package main

import (
	f "fmt"
	"math"
)

func main() {

	var x float64

	f.Print("Digite o valor de x (em radianos): ")
	f.Scan(&x)

	var soma float64 = 0

	for i := 0; i < 20; i++ {

		// calcula x^(2*i)
		pot := 1.0
		for i := 0; i < 2*i; i++ {
			pot *= x
		}

		// calcula (2*i)!
		fat := 1.0
		for j := 1; j <= 2*i; j++ {
			fat *= float64(j)
		}

		termo := pot / fat

		// alterna sinal
		if i%2 == 0 {
			soma += termo
		} else {
			soma -= termo
		}
	}

	// valor real
	real := math.Cos(x)

	f.Printf("Cos(x) pela série: %.6f\n", soma)
	f.Printf("Cos(x) real: %.6f\n", real)
	f.Printf("Diferença: %.6f\n", math.Abs(soma-real))

}
