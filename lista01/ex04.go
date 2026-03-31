// Sabendo-se que 100 kW de energia custam 70% do salário mínimo, escreva um algoritmo em Linguagem
// C que leia o valor do salário mínimo e a quantidade de kW gasta por uma residência. Calcule e
// imprima:
//  o valor em reais de cada kW;
//  o valor em reais a ser pago pelo consumo da residência;
//  o novo valor a ser pago pela residência com um desconto de 10%.
// Entrada
// O programa deve ler o valor do salário mínimo e a quantidade de kW gasta por uma residência. Ambos
// os valores são reais.
// Saída
// O programa deve imprimir três linhas contento o texto:
// Custo por kW: R$ x.xx
// Custo do consumo: R$ x.xx
// Custo com desconto: R$ x.xx

package main

import (
	"fmt"
)"

func main() {
	var salarioMinimo float64
	var qtdKW float64
	var custoPorKW float64
	var custoConsumo float64
	var custoDesconto float64

	fmt.Println("Digite o valor do salário mínimo")
	fmt.Scan(&salarioMinimo)

	fmt.Println("Digite a quantidade de KW consumida")
	fmt.Scan(&qtdKW)

	custoPorKW = (salarioMinimo * 0.7) / 100
	custoConsumo = custoPorKW * qtdKW
	custoDesconto = custoConsumo * 0.9

	fmt.Printf("custo por KW: R$ %.2f\n", custoPorKW)
	fmt.Printf("custo do consumo: R$ %.2f\n", custoConsumo)
	fmt.Printf("custo com desconto: R$ %.2f\n", custoDesconto)

}
