// 21. Faça um programa que leia um código numérico inteiros e um vetor de 10 posições de números reais. Se o
// código for zero, termine o programa. Se for 1, mostre o vetor na ordem direta. Se for 2, mostre o vetor na
// ordem inversa.

package main

import "fmt"

func main() {
	var vetor [10]float64
	var codigo int

	// 1. Leitura do vetor
	fmt.Println("Digite 10 números reais:")
	for i := 0; i < 10; i++ {
		fmt.Printf("Elemento %d: ", i+1)
		fmt.Scan(&vetor[i])
	}

	// 2. Leitura do código e lógica de exibição
	for {
		fmt.Print("\nDigite o código (0=Sair, 1=Direta, 2=Inversa): ")
		fmt.Scan(&codigo)

		if codigo == 0 {
			fmt.Println("Encerrando programa...")
			break
		}

		switch codigo {
		case 1:
			fmt.Println("Ordem Direta:", vetor)
		case 2:
			fmt.Print("Ordem Inversa: [")
			for i := 9; i >= 0; i-- {
				fmt.Printf("%v ", vetor[i])
			}
			fmt.Println("]")
		default:
			fmt.Println("Código inválido!")
		}
	}
}
