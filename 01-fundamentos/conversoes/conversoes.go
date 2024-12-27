package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	// fmt.Println(x / y) // error
	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	// Cuidado...
	fmt.Println("Teste " + string(97)) // converte o valor 97 para o caractere correspondente na tabela ASCII

	// int para string
	fmt.Println("Teste " + strconv.Itoa(97))

	// string para int
	num, _ := strconv.Atoi("123")
	fmt.Println(num)

	// string para boolean
	b, _ := strconv.ParseBool("true")
	if b {
		fmt.Println("Verdadeiro")
	}
}
