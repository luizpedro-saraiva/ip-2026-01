// 35. Faça um programa que leia um número inteiro positivo na base 10, calcule e imprima o seu equivalente na
// base 2.

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
		fmt.Println("Binário: 0")
		return
	}

	var binario []int

	for n > 0 {
		resto := n % 2
		binario = append(binario, resto)
		n = n / 2
	}

	fmt.Print("Binário: ")

	for i := len(binario) - 1; i >= 0; i-- {
		fmt.Print(binario[i])
	}

	fmt.Println()
}
