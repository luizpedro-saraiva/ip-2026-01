// 22. Faça um programa que simule um controle bancário. Para tanto, devem ser lidos os códigos de dez contas e os
// seus respectivos saldos. Os códigos devem ser armazenados em um vetor de números inteiros (não pode haver
// mais do que uma conta com o mesmo código) e os saldos devem ser armazenados em um vetor de números
// reais. O saldo deverá ser cadastrado na mesma posição do código. Por exemplo, se a conta 504 foi armazenada
// na 5a posição do vetor de códigos, o seu saldo deverá ficar na 5a posição do vetor de saldos. Depois de fazer a
// leitura dos valores, mostrar o seguinte menu na tela:
// 1. Efetuar depósito
// 2. Efetuar saque
// 3. Consultar o ativo bancário (ou seja, o somatório dos saldos de todos os clientes)
// 4. Finalizar o programa.
// - Para efetuar depósito deve-se solicitar o código da conta e o valor a ser depositado. Se a conta não estiver
// cadastrada, mostrar a mensagem Conta não encontrada e voltar ao menu. Se a conta existir, atualizar o seus
// saldo.
// - Para efetuar saque deve-se solicitar o código da conta e o valor a ser sacado. Se a conta não estiver
// cadastrada, mostrar a mensagem Conta não encontrada e voltar ao menu. Se a conta existir, verificar verificar
// se o seu saldo é suficiente para cobrir o saque. (Estamos supondo que a conta não pode ficar com o saldo
// negativo). Se o saldo for suficiente, realizar o saque e voltar ao menu. Caso contrário, mostrar a mensagem
// Saldo insuficiente e voltar ao menu.
// - Para consultar o ativo bancário deve-se somar o saldo de todas as contas do banco. Depois de mostrar esse
// valor, voltar ao menu.
// - O programa só termina quando for digitada a opção 4 – Finalizar programa.

package main

import "fmt"

func main() {
	var codigos [10]int
	var saldos [10]float64

	// 1. Cadastro das Contas
	fmt.Println("--- Cadastro de 10 Contas ---")
	for i := 0; i < 10; {
		var codigo int
		fmt.Printf("Código da conta %d: ", i+1)
		fmt.Scan(&codigo)

		// Verifica se o código já existe para evitar duplicatas
		duplicado := false
		for j := 0; j < i; j++ {
			if codigos[j] == codigo {
				duplicado = true
				break
			}
		}

		if duplicado {
			fmt.Println("Erro: Já existe uma conta com esse código. Tente outro.")
			continue
		}

		codigos[i] = codigo
		fmt.Printf("Saldo inicial da conta %d: ", codigo)
		fmt.Scan(&saldos[i])
		i++
	}

	// 2. Menu Principal
	for {
		fmt.Println("\n--- Menu Bancário ---")
		fmt.Println("1. Efetuar depósito")
		fmt.Println("2. Efetuar saque")
		fmt.Println("3. Consultar ativo bancário")
		fmt.Println("4. Finalizar programa")
		fmt.Print("Escolha uma opção: ")

		var opcao int
		fmt.Scan(&opcao)

		if opcao == 4 {
			fmt.Println("Programa finalizado.")
			break
		}

		switch opcao {
		case 1: // Depósito
			var conta int
			fmt.Print("Código da conta: ")
			fmt.Scan(&conta)

			indice := -1
			for i, v := range codigos {
				if v == conta {
					indice = i
					break
				}
			}

			if indice == -1 {
				fmt.Println("Conta não encontrada.")
			} else {
				var vlr float64
				fmt.Print("Valor do depósito: ")
				fmt.Scan(&vlr)
				saldos[indice] += vlr
				fmt.Printf("Depósito realizado! Novo saldo: R$ %.2f\n", saldos[indice])
			}

		case 2: // Saque
			var conta int
			fmt.Print("Código da conta: ")
			fmt.Scan(&conta)

			indice := -1
			for i, v := range codigos {
				if v == conta {
					indice = i
					break
				}
			}

			if indice == -1 {
				fmt.Println("Conta não encontrada.")
			} else {
				var vlr float64
				fmt.Print("Valor do saque: ")
				fmt.Scan(&vlr)
				if vlr > saldos[indice] {
					fmt.Println("Saldo insuficiente.")
				} else {
					saldos[indice] -= vlr
					fmt.Printf("Saque realizado! Novo saldo: R$ %.2f\n", saldos[indice])
				}
			}

		case 3: // Ativo Bancário
			ativo := 0.0
			for _, s := range saldos {
				ativo += s
			}
			fmt.Printf("Ativo bancário total (soma de todos os saldos): R$ %.2f\n", ativo)

		default:
			fmt.Println("Opção inválida.")
		}
	}
}
