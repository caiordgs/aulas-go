package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	c := cachorro{"Barto", "Spitz Alemão", 1}

	cEmJSON, erro := json.Marshal(c) // Transforma em array de bytes
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cEmJSON)
	fmt.Println(bytes.NewBuffer(cEmJSON)) // "Traduz" o array de bytes
}
