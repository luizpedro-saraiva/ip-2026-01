// Faça um programa que:
// - leia um número real X do teclado;
// - determine e imprima o seguinte somatório:
// S = X − X/1! + X/2! + X/3! + X/4! + ...
// usando os 20 primeiros termos da série.

package main

import f "fmt"

func main() {

	var x float64

	f.Println("Digite um número: ")
	f.Scan(&x)

	soma := x
	fatorial := 1.0

	// segundo termo (único negativo)
	soma -= x / fatorial

	// restantes positivos
	for i := 2; i < 20; i++ {
		fatorial *= float64(i)
		soma += x / fatorial
	}

	f.Printf("O somatório é: %f", soma)
}
