// 36. Faça um programa que leia um número inteiro positivo na base 10, calcule e imprima o seu equivalente na
// base 16.

package main

import "fmt"

func main() {

	var n int

	fmt.Print("Digite um número inteiro positivo: ")
	fmt.Scan(&n)

	if n < 0 {
		fmt.Println("Valor inválido.")
		return
	}

	if n == 0 {
		fmt.Println("Hexadecimal: 0")
		return
	}

	var resultado []string

	for n > 0 {
		resto := n % 16

		if resto < 10 {
			resultado = append(resultado, fmt.Sprintf("%d", resto))
		} else {
			// 10 -> A, 11 -> B ...
			letra := string('A' + (resto - 10))
			resultado = append(resultado, letra)
		}

		n = n / 16
	}

	fmt.Print("Hexadecimal: ")

	for i := len(resultado) - 1; i >= 0; i-- {
		fmt.Print(resultado[i])
	}

	fmt.Println()
}
