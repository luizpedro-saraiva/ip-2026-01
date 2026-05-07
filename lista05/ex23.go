// 23. Uma empresa possui ônibus com 48 lugares (24 nas janelas e 24 no corredor). Faça um programa que utilize
// dois vetores para controlar as poltronas ocupadas no corredor e na janela. Considere que zero representa
// poltrona desocupada e um representa poltrona ocupada.
// Janela [ 0 1 0 0 . . . 1 0 0 ]
// Corredor [ 0 0 0 1 . . . 1 0 0 ]
// Esse programa deve controlar a venda de passagens da seguinte maneira:
// - o cliente informa se deseja poltrona no corredor ou na janela e, depois, o programa deve informar quais
// poltronas estão disponíveis para a venda;
// - quando não existirem poltronas livres no corredor, nas janelas ou, ainda, quando o ônibus estiver
// completamente cheio, deve ser mostrada uma mensagem.

package main

import "fmt"

func main() {
	var janela [24]int // 0: livre, 1: ocupada
	var corredor [24]int
	var totalOcupado int

	for {
		// Verifica se o ônibus está lotado
		if totalOcupado == 48 {
			fmt.Println("\nÔnibus completamente lotado!")
			break
		}

		fmt.Println("\n--- Sistema de Vendas de Passagens ---")
		fmt.Println("1 - Janela")
		fmt.Println("2 - Corredor")
		fmt.Println("0 - Encerrar")
		fmt.Print("Escolha o tipo de poltrona: ")

		var opcao int
		fmt.Scan(&opcao)

		if opcao == 0 {
			break
		}

		var vetor *[24]int
		var tipo string

		if opcao == 1 {
			vetor = &janela
			tipo = "Janela"
		} else if opcao == 2 {
			vetor = &corredor
			tipo = "Corredor"
		} else {
			fmt.Println("Opção inválida!")
			continue
		}

		// Mostra poltronas disponíveis no tipo escolhido
		disponiveis := false
		fmt.Printf("Poltronas de %s disponíveis: ", tipo)
		for i := 0; i < 24; i++ {
			if vetor[i] == 0 {
				fmt.Printf("[%d] ", i+1)
				disponiveis = true
			}
		}
		fmt.Println()

		if !disponiveis {
			fmt.Printf("Não existem mais poltronas livres na %s.\n", tipo)
		} else {
			fmt.Print("Digite o número da poltrona desejada: ")
			var num int
			fmt.Scan(&num)

			indice := num - 1
			if indice < 0 || indice >= 24 || vetor[indice] == 1 {
				fmt.Println("Poltrona inválida ou já ocupada!")
			} else {
				vetor[indice] = 1
				totalOcupado++
				fmt.Println("Venda efetuada com sucesso!")
			}
		}
	}

	fmt.Println("Sistema encerrado.")
}
