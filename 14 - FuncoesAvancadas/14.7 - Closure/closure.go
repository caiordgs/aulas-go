package main

import "fmt"

func closure() func() {
	texto := "Dentro da função Closure" // funcaoNova puxa esse

	funcao := func() {
		fmt.Println(texto)
	}

	return funcao
}

func main() {
	texto := "Dentro da função Main"
	fmt.Println(texto)

	funcaoNova := closure()
	funcaoNova()
}
