// Escreva um programa que imprima uma tabela de conversão de graus Fahrenheit para graus Celsius.
// Dado um valor de temperatura F medida na escala Fahrenheit, seu valor equivalente C na escala Celsius é
// dado pela seguinte equação: C = (5(F-32)) / 9
// Entrada
// A entrada conterá várias linhas. A primeira delas contém o número n de temperaturas em Fahrenheit a
// serem convertidas para Celsius. Cada uma das n linhas seguintes contém um valor real (float) com a medida
// de uma temperatura em graus Fahrenheit.
// Saída
// O programa deve imprimir n linhas cada uma no seguinte formato x FAHRENHEIT EQUIVALE A y CELSIUS,
// onde x corresponde a um valor de temperatura em Fahrenheit e y corresponde ao valor equivalente em
// graus Celsius. Logo após a palavra CELSIUS em cada linha de saída deve ser impresso o caractere de
// quebra de linha. Os valores de x e y devem ser impressos com duas casas decimais.

package main

import (
	"fmt"
)

func main() {
	var Fahrenheit float64
	var Celsius float64
	var quantidadesTemperaturas int 
	
	fmt.Scan(&quantidadesTemperaturas)
	fmt.Println("Informe a quantidade de temperaturas que serão convertidas")	

	for i := 0; i < int(quantidadesTemperaturas); i++ {
		fmt.Printf("Informe a temperatura %d: ", i+1)
		fmt.Scan(&Fahrenheit)

		celsius = (5.0 * (fahrenheit - 32)) / 9

		fmt.Printf("%.2f FAHRENHEIT EQUIVALE A %.2f CELSIUS\n", fahrenheit, celsius)
}
