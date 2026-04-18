// Escreva um programa que imprima os índices de
// todos os elementos de uma matriz 10x10.

package main

import f "fmt"

func main() {

	var matriz [10][10]int
	f.Scan(&matriz)

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			f.Printf("Índice: [%d][%d]\n", i, j)

		}
	}

}
