// 10. A série de Fibonacci é formada pela sequência:
// 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, ...
// Escreva um programa que armazene em um vetor os primeiros 50 termos da série de Fibonacci. Após isso, o
// programa deve imprimir todos os valores armazenados.

package main

import f "fmt"

func main() {
	// Declarar um vetor para armazenar os 50 primeiros termos
	var fib [50]uint64

	// Definir os dois primeiros termos, que são a base da sequência
	fib[0] = 1
	fib[1] = 1

	// Calcular os próximos termos (cada um é a soma dos dois anteriores)
	for i := 2; i < 50; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}

	// Imprimir todos os valores armazenados
	f.Println("Os 50 primeiros termos da série de Fibonacci:")
	for i, valor := range fib {
		f.Printf("Termo %d: %d\n", i+1, valor)
	}
}
