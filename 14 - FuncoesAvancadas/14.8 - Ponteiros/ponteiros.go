package main

import "fmt"

func inverteSinal(numero int) int {
	return numero * -1
}

func inverteSinalComPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {
	numero := 20
	numeroInvertido := inverteSinal(numero)
	fmt.Println(numeroInvertido)

	novoNumero := 40
	fmt.Println(novoNumero)
	inverteSinalComPonteiro(&novoNumero) // ponteiro altera o valor da variavel novoNumero
	fmt.Println(novoNumero)
}
