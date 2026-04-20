// 32. Faça um programa que calcule e escreva a multiplicação de dois números N1 e
// N2 lidos do teclado. Obs.: a
// máquina que irá executar esse programa somente sabe adicionar e subtrair.

package main

import "fmt"

func main() {

	var n1, n2 int

	fmt.Print("Digite N1: ")
	fmt.Scan(&n1)

	fmt.Print("Digite N2: ")
	fmt.Scan(&n2)

	resultado := 0
	negativo := false

	// tratar sinal de N2 sem usar multiplicação
	if n2 < 0 {
		n2 = 0 - n2
		negativo = true
	}

	for i := 0; i < n2; i++ {
		resultado = resultado + n1
	}

	if negativo {
		resultado = 0 - resultado
	}

	fmt.Printf("Resultado: %d\n", resultado)
}
