// Fazer um algoritmo que calcule a média aritmética das três notas de um aluno e mostre, além do valor
// da média, uma mensagem de "APROVADO", caso a média seja igual ou superior a seis, ou a mensagem
// "REPROVADO", caso contrário.
// Entrada
// A entrada conterá uma linha com as três notas do aluno, separadas entre si por um caractere de espaço.
// Saída
// A saída deve conter duas linhas. A primeira linha deve conter uma frase com o seguinte formato:
// MEDIA = x, onde x é o valor da média entre as três notas do aluno, contendo duas casas decimais. A
// segunda linha contém uma das duas mensagens: APROVADO ou REPROVADO. Após o valor da média e
// após a mensagem, o programa deve imprimir o caractere de quebra de linha: ‘\n’.

package main

import (
	"fmt"
)

func main() {
	var x float32
	var y float32
	var z float32
	var media float32
	fmt.Scan(&x)
	fmt.Scan(&y)
	fmt.Scan(&z)
	media = (x + y + z) / 3
	fmt.Printf("MEDIA = %.2f\n", media)
	if media >= 6 {
		fmt.Println("APROVADO")
	} else {
		fmt.Println("REPROVADO")
	}
}
