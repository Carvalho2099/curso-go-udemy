package main

import "fmt"

func trocar(p1, p2 int) (segundo, primeiro int) {
	primeiro = p2
	segundo = p1
	return // retorno limpo
}

func main() {
	primeiro, segundo := trocar(1, 2)
	fmt.Println(primeiro, segundo)

	segundo, primeiro = trocar(3, 4)
	fmt.Println(primeiro, segundo)
}
