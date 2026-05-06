// 5. Faça um programa que leia um vetor de inteiros, de 10 posições. A seguir, encontre o menor elemento (X) do
// vetor. Imprima uma mensagem mostrando: “O menor elemento do vetor é”, X, “e sua posição dentro do vetor
// é:”,P. Assuma que os elementos informados no vetor são todos diferentes entre si.

package main

import (
	f "fmt"
)

func main() {
	// 1. Declaração do vetor de 10 posições
	var vetor [10]int
	var menorElemento int
	var posicaoMenor int

	// 2. Leitura dos dados do vetor
	f.Println("Digite 10 números inteiros diferentes:")
	for i := 0; i < 10; i++ {
		f.Printf("Posição %d: ", i)
		f.Scan(&vetor[i])
	}

	// 3. Inicializa o menorElemento com o primeiro elemento do vetor
	menorElemento = vetor[0]
	posicaoMenor = 0

	// 4. Encontra o menor elemento e sua posição
	for i := 1; i < 10; i++ {
		if vetor[i] < menorElemento {
			menorElemento = vetor[i]
			posicaoMenor = i
		}
	}

	// 5. Imprime o resultado
	f.Printf("\nO menor elemento do vetor é: %d\n", menorElemento)
	f.Printf("Sua posição dentro do vetor é: %d\n", posicaoMenor)
}
