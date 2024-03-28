package mai

import "fmt"

type employee struct {
	nome    string
	idade   int
	funcao  string
	salario int
}

func main() {

	employee1 := employee{nome: "Julia", idade: 22, funcao: "Developer", salario: 10000}
	employee2 := employee{nome: "Alexia", idade: 26, funcao: "ScrumMaster", salario: 8000}
	employee3 := employee{nome: "Gabi", idade: 25, funcao: "Support", salario: 8000}

	var employees [3]employee

	employees[0] = employee1
	employees[1] = employee2
	employees[2] = employee3

	fmt.Print(employees[1])

}
