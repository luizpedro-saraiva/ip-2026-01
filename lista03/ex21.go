// 21. Escreva um programa que calcule e imprima o valor de b elevado a n. O
//  usuário vai informar os valores de b e n. Assuma
// que o valor de n é maior do que 1 e o valor de b é maior
// ou igual a 2, ambos valores inteiros. Em sua solução
// não é permitido o uso da função pow.

package main

import f "fmt"

func main() {
	var b, n int

	f.Print("Digite o valor de b (base): ")
	f.Scan(&b)

	f.Print("Digite o valor de n (expoente): ")
	f.Scan(&n)

	resultado := 1

	for i := 0; i < n; i++ {
		resultado *= b
	}

	f.Printf("%d elevado a %d é: %d\n", b, n, resultado)
}
