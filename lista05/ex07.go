// 7. Escreva um programa que armazene os 100 primeiros números ímpares em um vetor. Imprima o vetor em
// seguida.

package main

import f "fmt"

func main() {

	var impares [100]int

	// 2. Preencher o vetor com os 100 primeiros números ímpares
	// O primeiro número ímpar é 1, e eles aumentam de 2 em 2.
	contador := 0
	for i := 1; contador < 100; i++ {
		if i%2 != 0 {
			impares[contador] = i
			contador++
		}
	}

	f.Println("Os 100 primeiros números ímpares:")
	f.Println(impares)

}
