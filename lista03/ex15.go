// 15. Seja a seguinte série:

// 1, 4, 9, 16, 25, 36, ...

// Escreva um programa que gere esta série até o N-ésimo termo. N será informado pelo usuário.

package main

import f "fmt"

func main() {

	var N int

	f.Print("Digite o valor de N: ")
	f.Scan(&N)

	if N <= 0 {
		f.Println("N deve ser maior que zero.")
		return
	}

	for i := 1; i <= N; i++ {
		f.Print(i*i, " ")
	}
}
