// Desenvolver um programa para calcular a conta de água para uma empresa de saneamento. O custo da
// água varia dependendo se o consumidor é residencial, comercial ou industrial. A regra para calcular a conta
// é:
// • Residencial: R$ 5, 00 de taxa mais R$ 0, 05 por metros cúbicos gastos;
// • Comercial: R$ 500, 00 para os primeiros 80 metros cúbicos gastos mais R$ 0, 25 por metros cúbicos
// gastos;
// • Industrial: R$ 800, 00 para os primeiros 100 metros cúbicos gastos mais R$ 0, 04 por metros cúbicos
// gastos;
// O programa deverá ler a conta do cliente, o consumo de água por metros cúbicos e o tipo de consumidor (
// residencial, comercial e industrial ). Como resultado, o programa deve imprimir a conta do cliente e o valor
// em Reais a ser pago pelo mesmo.
// Entrada

// O programa deverá ler uma linha na entrada contendo: a conta do cliente (um número inteiro), o con-
// sumo de água por metros cúbicos (float) e o tipo do consumidor (um caractere: ‘C’ - COMERCIAL, ‘I’ -

// INDUSTRIAL ou ‘R’ - RESIDENCIAL). Há um espaço entre os valores na linha de entrada
// Saída
// O programa deve imprimir duas linhas, contendo o seguinte:
// • CONTA = u, onde u é o código inteiro identificador da conta;
// • VALOR DA CONTA = v, onde v é o valor da conta com duas casas decimais, a ser pago pelo
// consumidor;
// Após o valor impresso em cada linha, o programa dev imprimir o caractere de quebra de linha; ‘\n’. Os
// valores de v,x e w devem conter duas casas decimais.

package main

import (
	"fmt"
)

func main() {

	var valor_Conta float64
	var conta_Cliente int
	var consumo_Agua float64
	var tipo_Cliente string

	fmt.Print("Informe o número da conta, valor do consumo de água e tipo de cliente: 'C' - Comercial, 'I' - Industrial ou 'R' - Residencial. ")

	fmt.Scan(&valor_Conta, &consumo_Agua, &tipo_Cliente)
	fmt.Scan(&conta_Cliente)

	if tipo_Cliente == "I" {
		valor_Conta = 800 + ((consumo_Agua - 100) * 0.04)
		fmt.Printf("Conta = %d\nVALOR DA CONTA = %.2f", conta_Cliente, valor_Conta)
	} else if tipo_Cliente == "C" {
		valor_Conta = 500 + ((consumo_Agua - 80) * 0.25)
		fmt.Printf("Conta = %d\nVALOR DA CONTA = %.2f", conta_Cliente, valor_Conta)
	} else if tipo_Cliente == "R" {
		valor_Conta = 5 + (consumo_Agua * 0.05)
		fmt.Printf("Conta = %d\nVALOR DA CONTA = %.2f", conta_Cliente, valor_Conta)
	} else {
		fmt.Print("Informe um tipo de cliente válido!")
	}

}
