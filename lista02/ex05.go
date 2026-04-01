// Escreva um programa que leia um número inteiro e diga se ele é ou não um número divisível por 5.

package main

import (
	"fmt"
)

func main() {

	var n int

	fmt.Println("Informe um valor")
	fmt.Scan(&n)

	if n%5 == 0 {
		fmt.Println("Divisível por 5")
	} else {
		fmt.Println("Não divisível por 5")
	}

}
