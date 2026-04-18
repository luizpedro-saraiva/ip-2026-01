// Escreva um programa que calcule o fatorial de um número
// inteiro N fornecido pelo usuário. Cuidado com
// valores inválidos!

package main

import f "fmt"

func main() {

	var n int

	f.Println("Digite um número inteiro para calcular o fatorial: ")
	f.Scan(&n)

	if n < 0 {
		f.Println("Fatorial não é definido para números negativos.")
		return
	} else if n == 0 || n == 1 {
		f.Printf("O fatorial de %d é 1.\n", n)
		return
	}

	fatorial := 1
	for i := 2; i <= n; i++ {
		fatorial *= i
	}

	f.Printf("O fatorial de %d é %d.\n", n, fatorial)
}
