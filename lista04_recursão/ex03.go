// Dado um array de inteiros e o seu número de elementos, inverta a
// posição dos seus elementos.

package main

import "fmt"

func main() {
	var n int

	fmt.Println("Digite o valor de n")
	fmt.Scan(&n)

	v := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&v[i])
	}

	inicio := 0
	fim := n - 1

	for inicio < fim {
		v[inicio], v[fim] = v[fim], v[inicio]
		inicio++
		fim--
	}

	for i := 0; i < n; i++ {
		fmt.Print(v[i], " ")
	}
}
