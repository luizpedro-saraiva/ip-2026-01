// 34. Faça um programa que leia dois números inteiros positivos N1 e N2, calcule e escreva o mínimo múltiplo
// comum para este par de números (N1,N2).

package main

import (
	"fmt"
)

func main() {

	var n1, n2 int

	fmt.Print("Digite N1: ")
	fmt.Scan(&n1)

	fmt.Print("Digite N2: ")
	fmt.Scan(&n2)

	if n1 <= 0 || n2 <= 0 {
		fmt.Println("Valores devem ser positivos.")
		return
	}

	a := n1
	b := n2

	// cálculo do MDC (Euclides)
	for b != 0 {
		resto := a % b
		a = b
		b = resto
	}

	mdc := a
	mmc := (n1 * n2) / mdc

	fmt.Printf("MMC(%d, %d) = %d\n", n1, n2, mmc)
}
