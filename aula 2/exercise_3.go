package main

import (
	"fmt"
)

func main() {
	var numero int
	fmt.Print("Digite um numero: ")
	fmt.Scan(&numero)

	if numero > 0 {
		fmt.Println("Positivo!")
	} else if numero < 0 {
		fmt.Println("Negativo!")
	} else if numero == 0 {
		fmt.Println("Zero!")
	}

}
