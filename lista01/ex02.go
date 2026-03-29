// Escrever um algoritmo que lê o público total de futebol, as percentagens de pessoas nas seguintes
// categorias: popular, geral, arquibancada e cadeiras e forneça a renda total do jogo. Sabe-se que o valor
// dos ingressos por categoria são dados pela tabela abaixo:
// Categoria Valor ingresso
// Popular R$ 1,00
// Geral R$ 5,00
// Arquibancada R$10,00
// Cadeiras R$ 20,00
// Entrada
// A entrada contém uma linha inicial com um valor inteiro informando o número de casos de testes que
// ocorrem nas linhas seguintes. Cada caso de teste seguinte é formado por uma linha contendo os seguintes
// valores, separados entre si por um espaço:
//  O número de pessoas que compraram ingresso para o jogo correspondente ao caso de teste.
//  A percentagem de pessoas que compraram ingresso na categoria Popular.
//  A percentagem de pessoas que compraram ingresso na categoria Geral.
//  A percentagem de pessoas que compraram ingresso na categoria Arquibancada.
//  A percentagem de pessoas que compraram ingresso na categoria Cadeiras.
// Saída
// O programa deve gerar uma linha para cada caso de teste na entrada, contendo a frase: A RENDA DO
// JOGO N. x E = y, onde x corresponde a ordem do caso de teste na entrada e y é um valor real com duas
// casas decimais que corresponde ao valor da renda total do jogo x.
package main

import (
	"fmt"
)

func main() {
	var casos int
	var pessoas int
	var pp, pg, pa, pc float64
	var total float64

	fmt.Println("informe o número de casos")
	fmt.Scan(&casos)

	for i := 1; i <= casos; i++ {
		fmt.Scan(&pessoas)
		fmt.Scan(&pp, &pg, &pa, &pc)
		total = ((float64(pessoas) * (pp / 100.0)) * 1) + ((float64(pessoas) * (pg / 100.0)) * 5) + ((float64(pessoas) * (pa / 100.0)) * 10) + ((float64(pessoas) * (pc / 100.0)) * 20)
		fmt.Printf("A RENDA DO JOGO N. %d E = %.2f", i, total)
	}
}
