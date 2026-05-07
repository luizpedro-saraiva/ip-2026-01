// 16. Dado um vetor com dados de 50 idades, elabore um programa que permita calcular a moda das idades. Obs.:
// Moda é o valor que tem maior incidência de repetições.

package main

import "fmt"

func main() {
	// Exemplo com vetor de 50 idades (pode ser preenchido via Scan também)
	idades := [50]int{
		20, 25, 20, 30, 25, 40, 20, 18, 20, 50,
		25, 20, 30, 30, 21, 22, 20, 25, 25, 25,
		18, 19, 20, 25, 30, 35, 40, 45, 50, 55,
		20, 20, 25, 25, 30, 30, 30, 18, 19, 20,
		21, 22, 23, 24, 25, 26, 27, 28, 29, 30,
	}

	// 1. Criar um mapa para contar as ocorrências
	contagem := make(map[int]int)
	for _, idade := range idades {
		contagem[idade]++
	}

	// 2. Encontrar a idade com a maior contagem
	moda := 0
	maiorFrequencia := 0

	for idade, frequencia := range contagem {
		if frequencia > maiorFrequencia {
			maiorFrequencia = frequencia
			moda = idade
		}
	}

	fmt.Printf("A moda é %d anos, repetindo-se %d vezes.\n", moda, maiorFrequencia)
}
