// 18. Escreva um programa que imprima os índices da diagonal principal de uma matriz 10x10.

package main

import f "fmt"

func main() {
	f.Println("Índices da diagonal principal:")

	for i := 0; i < 10; i++ {
		f.Printf("[%d][%d]\n", i, i)
	}
}
