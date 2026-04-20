// 24. Escreva um programa para gerar e escrever uma tabela com os valores do seno de um
// ângulo A em radianos, usando a série de Mac-Laurin apresentada a seguir:
// SenA = A − (A³/6) + (A⁵/120) − (A⁷/5040)

package main

import f "fmt"

func main() {

	var A, t1, t2, t3, t4, seno float64

	f.Println("A(rad)\t\tSen(A) aproximado")

	for i := 0; i <= 63; i++ {
		A = float64(i) / 10.0

		t1 = A
		t2 = (A * A * A) / 6.0
		t3 = (A * A * A * A * A) / 120.0
		t4 = (A * A * A * A * A * A * A) / 5040.0

		seno = t1 - t2 + t3 - t4

		f.Printf("%.1f\t\t%.6f\n", A, seno)
	}
}
