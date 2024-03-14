package main

import (
	"fmt"
)

func main() {
	var name string

	fmt.Println("Digite seu nome: ")

	fmt.Scan(&name)

	fmt.Println("Bem-vindo(a), " + name)
	fmt.Printf("Bem-vinda(a), %s !\n", name)
}
