package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} // Compiler counts the number of elements

	for i, numero := range numeros {
		fmt.Println(i, numero)
	}

	for _, num := range numeros {
		fmt.Println(num)
	}
}
