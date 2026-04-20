// 19. Escreva um programa que imprima os índices dos elementos de uma matriz 10x10 que se encontram acima da
// diagonal principal.

package main

import f "fmt"

func main() {
	f.Println("Índices acima da diagonal principal:")

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j > i {
				f.Printf("[%d][%d]\n", i, j)
			}
		}
	}
}
