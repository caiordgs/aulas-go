package main

import (
	"fmt"
)

func main() {
	// Array de 5 posições
	var array1 [5]string
	array1[0] = "Posição 1"
	fmt.Println(array1)

	// Array de 5 posições
	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	fmt.Println(array2)

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(slice)

	// Adicionar novos itens no slice
	slice = append(slice, 15, 16, 17, 18, 19, 20)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "Posição Alterada"
	fmt.Println(array2)
	fmt.Println(slice2)
}
