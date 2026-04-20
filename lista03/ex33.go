// 33. Faça um programa que:
// - leia dois números inteiros positivos (N1 e N2);
// - calcule e escreva para este par de números (N1 e N2), o quociente e o resto da divisão de N1 por N2. Obs.: a
// máquina que irá calcular o quociente e o resto desta divisão somente sabe adicionar e subtrair. Portanto, não
// são possíveis o uso das funções quociente e resto.
// Exemplo: N1 = 14 e N2 = 4.
// Quociente(14,4) = 3 e o Resto(14,4)=2
// Procedimento: 14 – 4 = 10, 10 – 4 = 6, 6 – 4 = 2, o resto é 2. Como ocorreram 3 subtrações sucessivas, o
// quociente é 3.

package main

import "fmt"

func main() {

	var n1, n2 int

	fmt.Print("Digite N1: ")
	fmt.Scan(&n1)

	fmt.Print("Digite N2: ")
	fmt.Scan(&n2)

	if n2 <= 0 {
		fmt.Println("Divisor inválido.")
		return
	}

	quociente := 0
	resto := n1

	for resto >= n2 {
		resto = resto - n2
		quociente = quociente + 1
	}

	fmt.Printf("Quociente: %d\n", quociente)
	fmt.Printf("Resto: %d\n", resto)
}
