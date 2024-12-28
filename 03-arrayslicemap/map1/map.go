package main

import "fmt"

func main() {
	//var aprovados map[int]string
	// Maps must be initialized
	aprovados := make(map[int]string)

	aprovados[12345678978] = "Maria"
	aprovados[12345678979] = "Pedro"
	aprovados[12345678980] = "Ana"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[12345678978])
	delete(aprovados, 12345678978)
	fmt.Println(aprovados[12345678978])
}
