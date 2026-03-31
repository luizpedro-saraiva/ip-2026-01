// Muitos países estão passando a utilizar o sistema métrico. Faça um programa para executar as seguintes
// conversões:
// • Ler uma temperatura em Fahrenheit e imprimir o equivalente em Celsius (C = (5F − 160)/9).
// • Ler uma quantidade de chuva dada em polegadas e imprimir o equivalente em milímetros (1 polegada
// = 25.4 mm).
// Entrada
// O programa deve ler dois valores na entrada: um valor em Fahrenheit e outro valor em polegadas.
// Ambos os valores são do tipo float. Cada valor ocorre em uma linha diferente na entrada.
// Saída
// O programa deve imprimir duas linhas. Aa primeira contém a frase: O VALOR EM CELSIUS = X,
// onde X é o valor de temperatura convertido de Fahrenheit para Celsius e deve ter duas casas decimais. A
// segunda linha deve conter a frase: A QUANTIDADE DE CHUVA E = Y, onde Y é o valor em milímetros
// correspondente ao valor em polegadas dado como entrada. Y é um valor real (float) e deve ter duas casas
// decimais. Logo após o valor de Y, o programa deve imprimir o caractere de quebra de linha ‘\n’.

package main

import (
	"fmt"
)

func main() {
	var fahrenheit float64
	var celsius float64
	var qtdChuva float64
	var qtdMilimetro float64

	fmt.Scan(&fahrenheit, &qtdChuva)
	fmt.Print("Informe quantidade de chuva em polegadas e a temperatura em Fahrenheit: ")

	celsius = ((5.0 * fahrenheit) - 160) / 9
	qtdMilimetro = qtdChuva * 25.4

	fmt.Printf("O VALOR EM CELSIUS = %.2f\nA QUANTIDADE DE CHUVA E = %.2f\n", celsius, qtdMilimetro)

}
