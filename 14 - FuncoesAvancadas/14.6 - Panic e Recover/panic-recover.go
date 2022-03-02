package main

import "fmt"

func recuperarExecucao() {
	if r := recover(); r != nil { // recover do panic, sem isso, programa para.
		fmt.Println("Execução recuperada!")
	}
}

func alunoAprovado(n1, n2 float64) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A MÉDIA É EXATAMENTE 6") // para o programa completamente
}

func main() {
	fmt.Println(alunoAprovado(7, 6))
	fmt.Println("Pós execução!")

}
