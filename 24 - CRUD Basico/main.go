package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	// CRUD - CREATE, READ, UPDATE, DELETE

	router := mux.NewRouter()
	fmt.Println("Conectado na porta 5000.")
	log.Fatal(http.ListenAndServe(":5000", router))
}
