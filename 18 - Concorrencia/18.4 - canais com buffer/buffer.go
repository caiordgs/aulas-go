package main

import "fmt"

func main() {
	canal := make(chan string, 2) // buffer valor 2, mais que 2 canais == deadlock.
	canal <- "Olá Mundo!"
	canal <- "Programando em Go!"

	mensagem := <-canal
	mensagem2 := <-canal

	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}
