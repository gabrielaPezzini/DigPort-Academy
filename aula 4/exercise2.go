package main

import "fmt"

func main() {

	const mensagem = "Hello"

	var nomes []string

	nomes = append(nomes, mensagem, "Nome1", "Nome2", "Nome3")

	segundoNome := nomes[2]

	fmt.Println("Slice de nomes", nomes)
	fmt.Println(mensagem+",", segundoNome)
}

func main() {
	const oi = "Hello"

	s := make([]string, 0, 4)

	s = append(s, oi, "Maria", "Ana", "Julia", "Rafael")

	v := s[2]

	fmt.Println("esse é meu slice:", s)
	fmt.Println(s[0], v)
}
