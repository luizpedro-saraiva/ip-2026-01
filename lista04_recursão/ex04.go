// Um problema típico em ciência da computação consiste
// em converter um número da sua forma decimal para a forma binária

package main

import "fmt"

func main() {
	var n int

	fmt.Println("Digite o valor de n")
	fmt.Scan(&n)

	if n == 0 {
		fmt.Println(0)
		return
	}

	bin := []int{}

	for n > 0 {
		resto := n % 2
		bin = append(bin, resto)
		n = n / 2
	}

	for i := len(bin) - 1; i >= 0; i-- {
		fmt.Print(bin[i])
	}
}
