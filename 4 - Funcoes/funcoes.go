package main

import "fmt"

func somar(n1, n2 int8) int8 {
	return n1 + n2
}

func calculosMat(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao

}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt

	}

	resultado := f("Texto da Função 1")
	fmt.Println(resultado)

	resultadoSoma, resultadoSubtracao := calculosMat(30, 40)
	fmt.Println(resultadoSoma, resultadoSubtracao)
}
