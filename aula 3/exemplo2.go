package main

import "fmt"

func main() {
	contador := 0
	for {
		fmt.Println(contador, "Este é um loop infinito.")

		contador++
		if contador >= 5 {
			break
		}
	}
	fmt.Println(contador, "é o fim do loop.")
}
