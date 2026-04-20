// Calcule a soma de todos os valores de um array de reais

package main

import "fmt"

func soma(v []float64, n int) float64 {
	if n == 0 {
		return 0
	}
	return v[n-1] + soma(v, n-1)
}

func main() {

	valores := []float64{1.5, 2.0, 3.5, 4.0}

	resultado := soma(valores, len(valores))

	fmt.Printf("Soma: %.2f\n", resultado)
}
