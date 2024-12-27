package main

import "fmt"

func main() {
	a := 3
	b := 2

	// Go não aceita operações aritméticas entre inteiros de tipos diferentes
	fmt.Println("Soma =>", a+b)
	fmt.Println("Subtração =>", a-b)
	fmt.Println("Divisão =>", a/b)
	fmt.Println("Multiplicação =>", a*b)
	fmt.Println("Módulo =>", a%b)

	// bitwise
	fmt.Println("E =>", a&b)  // 11 & 10 = 10
	fmt.Println("Ou =>", a|b) // 11 | 10 = 11
}
