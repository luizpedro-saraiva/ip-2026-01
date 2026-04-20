// 16. A série de Fetuccine é gerada da seguinte forma: os dois primeiros termos (inteiros) são fornecidos pelo
// usuário. A partir daí, os termos são gerados com a soma ou subtração dos dois termos anteriores, ou seja:
// Ai = Ai-1 + Ai-2, para i ímpar;
// Ai = Ai-1 – Ai-2, para i par.
// Crie um programa que imprima os N primeiros termos da série de Fetuccine, assumindo que o usuário digitará
// um N>=3.

package main

impot f "fmt"

func main() {

	var a1, a2, N int

	fmt.Print("Digite A1: ")
	fmt.Scan(&a1)

	fmt.Print("Digite A2: ")
	fmt.Scan(&a2)

	fmt.Print("Digite N (>=3): ")
	fmt.Scan(&N)

	if N < 3 {
		fmt.Println("N deve ser >= 3")
		return
	}

	// imprimir os dois primeiros termos
	fmt.Print(a1, " ", a2, " ")

	anterior := a1
	atual := a2

	for i := 3; i <= N; i++ {
		var proximo int

		if i%2 != 0 { // ímpar
			proximo = atual + anterior
		} else { // par
			proximo = atual - anterior
		}

		fmt.Print(proximo, " ")

		// atualizar os termos
		anterior = atual
		atual = proximo
	}
	