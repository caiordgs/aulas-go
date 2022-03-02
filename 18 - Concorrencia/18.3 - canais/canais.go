package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)
	go escrever("Olá Mundo!", canal)

	mensagem := <-canal
	fmt.Println(mensagem)

	for { // loop infinito, se não fechar o canal (linha 32), deadlock.
		mensagem, aberto := <-canal
		if !aberto {
			break // se continuar aberto, break do loop.
		}
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do Programa!")
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ { // só envia 5x o texto.
		canal <- texto
		time.Sleep(time.Second)
	}

	close(canal) // fecha o canal
}
