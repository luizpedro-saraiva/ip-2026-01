// Escreva um programa que receba 10 números inteiros, armazene-os em um vetor e mostre:
// a) os números pares digitados;
// b) a soma dos números pares digitados;
// c) os números ímpares digitados;
// d) a quantidade de números ímpares digitados.

package main

import (
	f "fmt"
)

func main() {
	// Define um vetor (array) para 10 inteiros
	var numeros [10]int
	var pares []int
	var impares []int
	somaPares := 0

	f.Println("Digite 10 números inteiros:")

	for i := 0; i < 10; i++ {
		f.Printf("%dº número: ", i+1)
		f.Scan(&numeros[i])

		if numeros[i]%2 == 0 {
			pares = append(pares, numeros[i])
			somaPares += numeros[i]
		} else {
			impares = append(impares, numeros[i])
		}
	}

	f.Println("\n--- Resultados ---")

	// a) Mostrar os números pares digitados
	f.Print("a) Números pares digitados: ")
	if len(pares) == 0 {
		f.Println("Nenhum par digitado.")
	} else {
		f.Println(pares)
	}

	// b) Mostrar a soma dos números pares digitados
	f.Printf("b) Soma dos números pares: %d\n", somaPares)

	// c) Mostrar os números ímpares digitados
	f.Print("c) Números ímpares digitados: ")
	if len(impares) == 0 {
		f.Println("Nenhum ímpar digitado.")
	} else {
		f.Println(impares)
	}

	f.Printf("d) Quantidade de números ímpares: %d\n", len(impares))
}
