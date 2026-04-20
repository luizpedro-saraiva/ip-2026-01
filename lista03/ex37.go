// 37. Faça um programa que leia um número inteiro positivo na base 8, calcule e imprima o seu equivalente na base
// 10.

package main

import "fmt"

func main() {

	var octal int

	fmt.Print("Digite um número em base 8: ")
	fmt.Scan(&octal)

	if octal < 0 {
		fmt.Println("Valor inválido.")
		return
	}

	decimal := 0
	potencia := 1 // 8^0

	for octal > 0 {

		digito := octal % 10

		// valida se é realmente octal
		if digito >= 8 {
			fmt.Println("Número inválido na base 8.")
			return
		}

		decimal += digito * potencia

		potencia *= 8
		octal /= 10
	}

	fmt.Printf("Equivalente em base 10: %d\n", decimal)
}
