// Faça um programa que calcule o valor de H, sendo que H é determinado pela série:
// H = 1/1 + 3/2 + 5/3 + 7/4 + … + 99/50

package main

import f "fmt"

func main() {

	var H float64 = 0.0

	for i := 1; i <= 50; i++ {
		H += float64(2*i-1) / float64(i)
	}

	f.Printf("Valor de H: %.6f\n", H)

}
