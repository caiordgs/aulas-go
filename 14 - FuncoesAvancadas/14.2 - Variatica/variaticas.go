package main

import "fmt"

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}

	return total

}

func main() {
	totalDaSoma := soma(1, 2, 3, 4, 5, 28, 255, 82, 65, 41)
	fmt.Println(totalDaSoma)
}
