package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint16
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	u.nome = "Caio"
	u.idade = 28
	fmt.Println(u)

	enderecoEx := endereco{"Doutor José Maria", 11050}

	usuario2 := usuario{"Caio", 28, enderecoEx}
	fmt.Println(usuario2)
}
