package main

import "fmt"

func main() {
	numero := -8

	if numero > 15 {
		fmt.Println("É maior que 15.")

	} else {
		fmt.Println("Menor ou igual a 15.")
	}

	if outroNumero := numero; outroNumero > 0 { // essa variável outroNumero só existe dentro deste 'if/else'
		fmt.Println("Número maior que zero.")
	} else if outroNumero >= -10 {
		fmt.Println("Número entre -10 e zero.")
	} else {
		fmt.Println("Número -11 ou menor.")
	}
}
