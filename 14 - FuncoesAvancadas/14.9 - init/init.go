package main

import "fmt"

var n int

func init() { // vai ser executada antes de qualquer coisa, não importa a ordem.
	fmt.Println("Executando a função init")
	n = 10
}

func main() {
	fmt.Println("Função main sendo executada.")
	fmt.Println(n)
}
