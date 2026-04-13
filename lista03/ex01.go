// Escreva um programa que calcule potências. O usuário deve digitar a base e o expoente, e o programa deve
// apresentar o resultado (sem usar o comando pow). Assuma que o usuário irá digitar valores positivos.

package main

import f "fmt"

func main() {

	var base, expoente int
	var resultado int64 = 1

	// Soliticitar a base
	f.Println("Digite o valor da base: ")
	f.Scan(&base)

	// Solicitar o expoente
	f.Println("Digite o valor do expoente: ")
	f.Scan(&expoente)

	if expoente == 0 {
		resultado = 1

	} else {

		for i := 0; i < expoente; i++ {
			resultado *= int64(base)
		}
	}

	f.Printf("%d elevado a %d é: %d\n", base, expoente, resultado)
}
