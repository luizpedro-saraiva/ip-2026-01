// 31. Faça um programa que calcule e escreva o número de grãos de milho que
// se pode colocar em um tabuleiro e
// xadrez, colocando 1 no primeiro quadro e nos quadros seguintes o dobro
// do quadro anterior. São 64 quadros no
// total.

package main

import (
	"fmt"
	"math/big"
)

func main() {

	grao := big.NewInt(1)
	total := big.NewInt(0)

	for i := 0; i < 64; i++ {
		total.Add(total, grao)
		grao.Mul(grao, big.NewInt(2))
	}

	fmt.Println("Total de grãos:", total)
}
