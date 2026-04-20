// 30. Faça um programa que calcule o volume de uma esfera em função do raio R. O raio
// deverá variar de 0 a 20 cm
// de 0,5 em 0,5 cm. O volume é dado por:
// Volume = 4/3 * ∏ * R3

package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("Raio (cm)\tVolume (cm³)")

	for i := 0; i <= 40; i++ {
		R := float64(i) * 0.5

		volume := (4.0 / 3.0) * math.Pi * R * R * R

		fmt.Printf("%.1f\t\t%.4f\n", R, volume)
	}
}
