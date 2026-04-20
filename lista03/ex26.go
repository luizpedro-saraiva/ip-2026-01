// 26. Faça um programa que calcule e escreva a soma dos 20 primeiros termos da série:
// 100/0! + 99/1! + 98/2! + 97/3! + ...

package main

import "fmt"

func main() {

	var soma float64
	var fatorial float64 = 1.0

	for k := 0; k < 20; k++ {

		if k > 0 {
			fatorial *= float64(k)
		}

		numerador := float64(100 - k)
		termo := numerador / fatorial

		soma += termo
	}

	fmt.Printf("Soma dos 20 termos: %.6f\n", soma)
}
