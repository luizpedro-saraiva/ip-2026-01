// Faça um programa que leia um número do tipo float e imprima sua raiz quadrada caso o mesmo seja
// positivo ou nulo. Caso ele seja negativo, mostre o seu quadrado.

package main

import (
	"fmt"
	"math"
)

func main() {

	var n float64

	fmt.Println("Informe o número")
	fmt.Scan(&n)

	if n >= 0 {
		fmt.Println("Raiz quadrada:", math.Sqrt(n))
	} else {
		fmt.Println("Quadrado:", n*n)
	}

}
