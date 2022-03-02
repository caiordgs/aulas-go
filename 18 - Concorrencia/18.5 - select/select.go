package main

import (
	"fmt"
	"time"
)

func main() {
	canal1, canal2 := make(chan string), make(chan string) // criados os canais

	go func() { // função pra enviar pro canal a cada 0,5 seg
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "Canal 1"
		}
	}()

	go func() { // função pra enviar pro canal a cada 2,1 seg
		for {
			time.Sleep(time.Millisecond * 2005)
			canal2 <- "Canal 2"
		}
	}()

	for { // loop infinito. casos separados para uma mensagem não ter que esperar a outra chegar.
		select {
		case mensagemCanal1 := <-canal1:
			fmt.Println(mensagemCanal1)
		case mensagemCanal2 := <-canal2:
			fmt.Println(mensagemCanal2)
		}
	}
}
