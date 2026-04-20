// Dados dois números inteiros x e n, calcula o valor de x elevado a n

package main

import "fmt"

func potencia(x, n int) int {
	if n == 0 {
		return 1
	}
	return x * potencia(x, n-1)
}

func main() {

	var x, n int

	fmt.Print("Digite x: ")
	fmt.Scan(&x)

	fmt.Print("Digite n: ")
	fmt.Scan(&n)

	resultado := potencia(x, n)

	fmt.Printf("%d elevado a %d = %d\n", x, n, resultado)
}
