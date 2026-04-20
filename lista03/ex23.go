// 23. Faça um programa que calcule o resultado da série abaixo, considerando os N primeiros termos.
// N deve ser fornecido pelo usuário. Cuidado com valores inválidos.
// 1000/1 − 997/2 + 994/3 − 991/4 + ...

package main

import f "fmt"

func main() {
	var N int

	f.Print("Digite o valor de N: ")
	_, err := f.Scan(&N)

	if err != nil {
		f.Println("Entrada inválida.")
		return
	}

	if N <= 0 {
		f.Println("N deve ser maior que zero.")
		return
	}

	var soma float64 = 0

	for i := 1; i <= N; i++ {
		numerador := float64(1000 - 3*(i-1))
		termo := numerador / float64(i)

		if i%2 == 0 {
			soma -= termo
		} else {
			soma += termo
		}
	}

	f.Printf("Resultado: %.6f\n", soma)
}
