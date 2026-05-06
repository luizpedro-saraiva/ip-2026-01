// 12. Faça um programa que:
// - leia um conjunto de valores inteiros correspondentes a 15 notas dos alunos de uma turma. As notas variam de
// 0 a 10;
// - calcule a frequência absoluta e a frequência relativa de cada nota;
// - imprima uma tabela contendo os valores das notas (0 a 10) e suas respectivas frequências absoluta e relativa.
// Observações:
// - Frequência absoluta de um nota é o número de vezes em que ela aparece no conjunto de dados;
// - Frequência relativa é a frequência absoluta dividida pelo número total de dados.

package main

import "fmt"

func main() {
	var notas [15]int
	var freqAbsoluta [11]int
	totalAlunos := 15

	fmt.Println("Digite as 15 notas (entre 0 e 10):")

	for i := 0; i < totalAlunos; i++ {
		fmt.Printf("Nota do aluno %d: ", i+1)
		fmt.Scan(&notas[i])

		// Validação simples para garantir que a nota está no intervalo 0-10
		if notas[i] >= 0 && notas[i] <= 10 {
			freqAbsoluta[notas[i]]++
		} else {
			fmt.Println("Nota inválida! Por favor, digite um valor entre 0 e 10.")
			i--
		}
	}

	// 2. Impressão da tabela de frequências
	fmt.Println("\nNota | Freq. Absoluta | Freq. Relativa")
	fmt.Println("---------------------------------------")

	for nota := 0; nota <= 10; nota++ {
		fAbs := freqAbsoluta[nota]
		// Cálculo da frequência relativa (frequência absoluta / total)
		fRel := float64(fAbs) / float64(totalAlunos)

		fmt.Printf("%4d | %14d | %14.2f%%\n", nota, fAbs, fRel*100)
	}
}
