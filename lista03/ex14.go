// 14. Faça um programa que imprima todos os números primos existentes entre N1 e N2, em que N1 e N2 são
// números naturais fornecidos pelo usuário.

package main

import f "fmt"

func main() {

	var n1, n2 int

	f.Print("Digite N1: ")
	f.Scan(&n1)
	f.Print("Digite N2: ")
	f.Scan(&n2)

	// garantir ordem correta
	if n1 > n2 {
		n1, n2 = n2, n1
	}

	for i := n1; i <= n2; i++ {
		if i < 2 {
			continue
		}

		primo := true

		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				primo = false
				break
			}
		}

		if primo {
			f.Print(i, " ")
		}
	}
}
