// Faça um programa que preencha um vetor com 10 números inteiros. Calcule e mostre os números superiores a
// 50 e suas respectivas posições. O programa deverá mostrar uma mensagem se não existir nenhum número
// nessa condição.
package main

import (
	f "fmt"
)

func main() {

	var numeros [10]int
	var encontrado bool = false

	// 2. Preenchimento do vetor
	f.Println("Digite 10 números inteiros:")
	for i := 0; i < 10; i++ {
		f.Printf("Número %d: ", i+1)
		f.Scan(&numeros[i])
	}

	f.Println("\n--- Resultados ---")

	// 3. Cálculo e exibição dos números superiores a 50 e suas posições
	for i := 0; i < 10; i++ {
		if numeros[i] > 50 {
			// i é o índice (0-9), i+1 é a posição (1-10)
			f.Printf("Número %d encontrado na posição %d (índice %d)\n", numeros[i], i+1, i)
			encontrado = true
		}
	}

	if !encontrado {
		f.Println("Nenhum número superior a 50 foi informado.")
	}
}
