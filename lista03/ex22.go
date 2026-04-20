// Faça um programa que calcule e imprima a seguinte soma:
// S = (37×38)/1 + (36×37)/2 + (35×36)/3 + ... + (1×2)/37

package main

import f "fmt"

func main() {

	var S float64

	for i := 1; i <= 37; i++ {
		S += (float64(38-i) * float64(39-i)) / float64(i)
	}

	f.Printf("O valor da soma S é: %.2f\n", S)
}
