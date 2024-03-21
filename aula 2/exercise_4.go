package main

import "fmt"

func main() {

	fmt.Println("Contagem regressiva:")
	i := 10
	for i != 0 {
		fmt.Println(i)
		i = i - 1
		fmt.Println(i)
	}

}
