// 39. Num frigorífico existem 90 bois. Cada boi traz preso no seu pescoço um
// cartão contendo um número de
// identificação e seu peso. Implementar um programa que escreva o número e o
// peso do boi mais gordo e do boi
// mais magro (não é necessário armazenar os dados de todos os bois).

package main

import "fmt"

func main() {

	var numero, peso int

	var maiorPeso, menorPeso int
	var numMaior, numMenor int

	for i := 1; i <= 90; i++ {

		fmt.Printf("Boi %d - Número: ", i)
		fmt.Scan(&numero)

		fmt.Printf("Boi %d - Peso: ", i)
		fmt.Scan(&peso)

		if i == 1 {
			maiorPeso = peso
			menorPeso = peso
			numMaior = numero
			numMenor = numero
		} else {

			if peso > maiorPeso {
				maiorPeso = peso
				numMaior = numero
			}

			if peso < menorPeso {
				menorPeso = peso
				numMenor = numero
			}
		}
	}

	fmt.Println("Boi mais gordo:")
	fmt.Printf("Número: %d | Peso: %d\n", numMaior, maiorPeso)

	fmt.Println("Boi mais magro:")
	fmt.Printf("Número: %d | Peso: %d\n", numMenor, menorPeso)
}
