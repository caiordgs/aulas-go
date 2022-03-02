package main

import (
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
	cEmJSON := `{"nome":"Barto","raca":"Spitz Alemão","idade":1}`

	var c cachorro

	if erro := json.Unmarshal([]byte(cEmJSON), &c); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(c)
}
