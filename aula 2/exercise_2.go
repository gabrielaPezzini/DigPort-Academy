package main

import (
	"fmt"
)

func main() {
	var firstValue = 10

	var secondValue int
	fmt.Print("Digite um número: ")
	fmt.Scan(&secondValue)

	soma := firstValue + secondValue

	fmt.Println("Resultado: ", soma)
}
