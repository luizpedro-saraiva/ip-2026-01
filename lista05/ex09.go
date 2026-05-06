// 9. Escreva um programa que receba a altura de 10 atletas. Esse programa deve imprimir as alturas daqueles
// atletas que têm altura maior do que a média.

package main

import f "fmt"

func main() {
	var alturas [10]float64
	var soma, media float64

	// 1. Receber as 10 alturas e somar para a média
	f.Println("Digite a altura de 10 atletas:")
	for i := 0; i < 10; i++ {
		f.Printf("Atleta %d: ", i+1)
		f.Scan(&alturas[i])
		soma += alturas[i]
	}

	// 2. Calcular a média
	media = soma / 10
	f.Printf("\nMédia de altura: %.2f\n", media)

	// 3. Imprimir apenas as alturas maiores que a média
	f.Println("Atletas com altura acima da média:")
	for _, altura := range alturas {
		if altura > media {
			f.Printf("%.2f ", altura)
		}
	}
	f.Println()
}
