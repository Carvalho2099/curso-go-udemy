package main

import "fmt"

var f = func() {
	fmt.Println("Função anônima!")
}

var soma = func(a, b int) int {
	return a + b
}

func main() {
	f()

	fmt.Println(soma(2, 3))

	sub := func(a, b int) int {
		return a - b
	}

	fmt.Println(sub(2, 3))
}
