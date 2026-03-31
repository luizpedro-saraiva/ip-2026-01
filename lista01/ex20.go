// Fazer um programa que leia um valor de tempo expresso em horas, minutos e segundos e que converta
// esse tempo para um valor em segundos.
// Entrada
// O programa deve ler três linhas na entrada. A primeira contém um valor em horas, a segunda, contém
// um valor em minutos e a terceira, contém um valor em segundos. Os valores são todos números inteiros.
// Saída
// O programa deve imprimir uma linha contendo a frase: O TEMPO EM SEGUNDOS E = X, onde X é
// o valor do tempo convertido em segundos. Após o valor do tempo em segundos, o programa deve imprimir
// um caractere de quebra de linha: ‘\n’.

package main

import (
	"fmt"
)

func main() {
	var horas, minutos, segundos int
	var total int

	fmt.Print("Digite as horas: ")
	fmt.Scan(&horas)

	fmt.Print("Digite os minutos: ")
	fmt.Scan(&minutos)

	fmt.Print("Digite os segundos: ")
	fmt.Scan(&segundos)

	total = (horas * 3600) + (minutos * 60) + segundos

	fmt.Printf("O TEMPO EM SEGUNDOS E = %d\n", total)
}
