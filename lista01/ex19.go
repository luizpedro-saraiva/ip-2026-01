// Faça um programa que leia um valor n, inteiro e positivo, calcule e mostre a seguinte soma:

// Entrada
// O programa deve ler um número inteiro positivo e maior que 1.
// Saída
// O programa deve apresentar uma linha contendo o valor final do somatório com 6 casas decimais. Caso
// o número lido não atenda as especificações da entrada, o programa deve apresentar a mensagem: "Numero
// invalido!".
// Observações
// Use precisão dupla para o cálculo de S.

package main

import (
	"fmt"
)

func main() {
	var a1, r, n int
	var somaFinal int
	var termo int

	fmt.Scan(&a1, &r, &n)

	termo = a1

	if n > 0 {
		for i := 1; i <= n; i++ {
			somaFinal += termo
			termo += r
		}
		fmt.Println(somaFinal)
	} else {
		fmt.Println("Valor invalido!")
	}
}
