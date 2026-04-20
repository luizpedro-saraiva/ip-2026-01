// 29. Escreva um programa que calcule e imprima o valor do somatório de todos os
// números inteiros de 1 a N, onde
// N é um número inteiro positivo fornecido pelo usuário.

package main

import "fmt"

func main() {

	var N int

	fmt.Print("Digite o valor de N: ")
	_, err := fmt.Scan(&N)

	if err != nil || N <= 0 {
		fmt.Println("Valor inválido.")
		return
	}

	soma := 0

	for i := 1; i <= N; i++ {
		soma += i
	}

	fmt.Printf("Somatório de 1 até %d: %d\n", N, soma)
}
