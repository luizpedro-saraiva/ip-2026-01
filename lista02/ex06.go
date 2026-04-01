// Crie um programa para determinar se um número inteiro A é divisível por outro número B. Os
// valores devem ser fornecidos pelo usuário.

package main

import (
	"fmt"
)

func main() {

	var A, B int

	fmt.Println("Digite o valor de A")
	fmt.Scan(&A)

	fmt.Println("Digite o valor de B")
	fmt.Scan(&B)

	if B == 0 {
		fmt.Println("Não é possível verificar divisibilidade por zero")
		return
	}

	if A%B == 0 {
		fmt.Println("A é divisível por B")
	} else {
		fmt.Println("A não é divisível por B")
	}
}
