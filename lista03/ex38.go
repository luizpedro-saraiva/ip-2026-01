// 38. O cálculo dos dígitos verificadores do CPF ocorre da seguinte maneira:
// 1o) o primeiro dígito verificador é calculado a partir dos 9 números iniciais, multiplicando-se cada um, da
// direita para a esquerda, por números crescentes a partir do número 2, conforme o exemplo seguinte:
// CPF → 1 1 1 4 4 4 7 7 7
// Nros. Crescentes → 10 9 8 7 6 5 4 3 2
// Resultado → 10 9 8 28 24 20 28 21 14
// Em seguida, deve ser calculada a soma dos resultados das multiplicações: 10 + 9 + 8 + 28 + 24 + 20 + 28 + 21
// + 14 = 162
// Logo após, é preciso realizar a divisão inteira da soma por 11: 162/11 = 14 com resto 8
// O primeiro digito verificador do CPF é, então, obtido com base na seguinte condição:

// - Resto da divisão menor que 2, então o dígito será igual a 0 (Zero)
// - Resto da divisão maior ou igual a 2, então o dígito será igual a 11 menos o resto (11-8=3)

// Assim, o CPF com o seu primeiro dígito verificador já pode ser escrito: 111.444.777-3X
// 2o) Para obter o segundo dígito é preciso repetir os cálculos realizados anteriormente, considerando, agora, um
// número de 10 dígitos, onde o décimo dígito é o primeiro dígito verificador já calculado:

// CPF → 1 1 1 4 4 4 7 7 7 3
// Nros. Crescentes → 11 10 9 8 7 6 5 4 3 2
// Resultado → 11 10 9 32 28 24 35 28 21 6

// 11 + 10 + 9 +32 + 28 + 24 + 35 + 28 +21 + 6 = 204/11 = 18 com resto 6
// 11 – 6 = 5
// Dessa forma, o CPF será 111.444.777-35
// Diante do exposto, implemente um programa que realize a leitura de um CPF e seus dígitos verificadores e
// informe se o mesmo é válido ou não.

package main

import "fmt"

func main() {

	var cpf string
	fmt.Print("Digite o CPF (11 dígitos): ")
	fmt.Scan(&cpf)

	if len(cpf) != 11 {
		fmt.Println("CPF inválido")
		return
	}

	var n [11]int

	for i := 0; i < 11; i++ {
		n[i] = int(cpf[i] - '0')
	}

	// 1º dígito
	soma := 0
	peso := 10

	for i := 0; i < 9; i++ {
		soma += n[i] * peso
		peso--
	}

	resto := soma % 11
	d1 := 0
	if resto >= 2 {
		d1 = 11 - resto
	}

	// 2º dígito
	soma = 0
	peso = 11

	for i := 0; i < 10; i++ {
		soma += n[i] * peso
		peso--
	}

	resto = soma % 11
	d2 := 0
	if resto >= 2 {
		d2 = 11 - resto
	}

	// verificação
	if d1 == n[9] && d2 == n[10] {
		fmt.Println("CPF válido")
	} else {
		fmt.Println("CPF inválido")
	}
}
