// 40. Uma Companhia de teatro planeja dar uma série de espetáculos. A direção calcula que, a R$ 6,00 o ingresso,
// serão vendidos 130 ingressos e as despesas montarão em R$300,00. A uma diminuição de R$ 0,60 no preço
// dos ingressos espera-se que haja um aumento de 30 ingressos vendidos. Fazer um programa que escreva uma
// tabela de valores do lucro esperado em função do preço do ingresso, fazendo-se variar este preço de R$ 6,00 a
// R$ 1,00 de R$ 0,60 em R$ 0,60. Escreva ainda o lucro máximo esperado, o preço e o número de ingressos
// correspondentes.

package main

import "fmt"

func main() {

	preco := 6.0
	ingressos := 130

	var lucro, lucroMax float64
	var melhorPreco float64
	var melhorIngressos int

	fmt.Println("Preco\tIngressos\tLucro")

	for preco >= 1.0 {

		lucro = preco*float64(ingressos) - 300

		fmt.Printf("%.2f\t%d\t\t%.2f\n", preco, ingressos, lucro)

		if lucro > lucroMax || preco == 6.0 {
			lucroMax = lucro
			melhorPreco = preco
			melhorIngressos = ingressos
		}

		preco -= 0.6
		ingressos += 30
	}

	fmt.Println("\nLucro máximo:")
	fmt.Printf("Lucro: %.2f\n", lucroMax)
	fmt.Printf("Preço: %.2f\n", melhorPreco)
	fmt.Printf("Ingressos: %d\n", melhorIngressos)
}
