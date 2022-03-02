package main

import "fmt"

func main() {

	// Aritméticos
	soma := 4 + 5
	subtracao := 27 - 32
	divisao := 8 / 6
	multiplicacao := 7 * 8
	restoDaDivisao := 5 % 4

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	// Relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 != 2)

	// Lógicos
	fmt.Println("------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)

	// Unários
	fmt.Println("------------")
	numero := 10 // valor 10
	numero++     // adiciona 1
	numero += 25 // numero = numero + 25
	numero--     // subtrai 1
	numero -= 15 // numero = numero - 15
	numero *= 3  // numero = numero * 3
	numero /= 2  // numero = numero / 2
	fmt.Println(numero)

}
