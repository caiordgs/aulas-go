package auxiliar

import (
	"fmt"
)

// Escrever registra mensagem
func Escrever() {
	fmt.Println("Escrevendo do arquivo do pacote auxiliar")
	//Não está com a primeira letra maiúscula mas está no mesmo pacote, por isso chama corretamente.
	escrever2()
}
