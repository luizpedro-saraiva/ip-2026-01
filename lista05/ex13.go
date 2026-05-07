// 13. Uma grande firma deseja saber quais os três empregados mais recentes. Faça um programa para ler um número
// indeterminado de informações (máximo de 100), contendo o número do empregado e o número de meses de
// trabalho deste empregado. Imprima os três empregados que entraram para trabalhar mais recentemente na
// firma. Obs.: a última informação contém os dois números iguais a zero. Assuma que não existem dois ou mais
// empregados admitidos no mesmo mês.

package main

import "fmt"

type Empregado struct {
	ID    int
	Meses int
}

func main() {

	e1 := Empregado{ID: -1, Meses: 999999}
	e2 := Empregado{ID: -1, Meses: 999999}
	e3 := Empregado{ID: -1, Meses: 999999}

	count := 0
	fmt.Println("Digite o ID e os meses (0 0 para sair):")

	for count < 100 {
		var id, meses int
		fmt.Scan(&id, &meses)

		if id == 0 && meses == 0 {
			break
		}

		novo := Empregado{ID: id, Meses: meses}

		if novo.Meses < e1.Meses {
			// Novo é o mais recente de todos: empurra o 1º e o 2º para baixo
			e3 = e2
			e2 = e1
			e1 = novo
		} else if novo.Meses < e2.Meses {
			// Novo é o segundo mais recente: empurra o 2º para baixo
			e3 = e2
			e2 = novo
		} else if novo.Meses < e3.Meses {
			// Novo é apenas o terceiro mais recente
			e3 = novo
		}

		count++
	}

	fmt.Println("\n--- 3 Empregados mais recentes ---")
	if e1.ID != -1 {
		fmt.Printf("1º - ID: %d (%d meses)\n", e1.ID, e1.Meses)
	}
	if e2.ID != -1 {
		fmt.Printf("2º - ID: %d (%d meses)\n", e2.ID, e2.Meses)
	}
	if e3.ID != -1 {
		fmt.Printf("3º - ID: %d (%d meses)\n", e3.ID, e3.Meses)
	}
}
