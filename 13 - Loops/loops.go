package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	for i < 10 {
		i++
		fmt.Println(i)
		time.Sleep(time.Second)
	}

	for j := 0; j < 10; j++ { // j criado aqui, não existe fora do 'for'
		fmt.Println(j)
		time.Sleep(time.Second)
	}

	nomes := [3]string{"Caio", "Sofia", "Bartô"}

	for _, nome := range nomes {
		fmt.Println(nome)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Caio",
		"sobrenome": "Rodrigues",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}
}
