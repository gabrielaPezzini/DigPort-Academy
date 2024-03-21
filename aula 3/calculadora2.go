package main

import "fmt"

func main() {
	fmt.Println("Digite um dos valores abaixo:")
	fmt.Println("1 - Soma")
	fmt.Println("2 - Subtração")
	fmt.Println("3 - Multiplicação")
	fmt.Println("4 - Divisão")

	var operacao int
	fmt.Scan(&operacao)

	if operacao == 1 {
		fmt.Println("Resultado: ", 1+1)
	} else if operacao == 2 {
		fmt.Println("Resultado: ", 1-1)
	} else {
		fmt.Println("")
	}

	switch operacao {
	case 1:
		fmt.Println("Resultado ", num1+num2)
	case 2:
		fmt.Println("Resultado ", num1-num2)
	case 3:
		fmt.Println("Resultado ", num1*num2)
	case 4:
		fmt.Println("Resultado ", num1/num2)
	default:
		fmt.Println("Opção Inválida")
	}
}
