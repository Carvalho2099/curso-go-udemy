package main

import (
	"fmt"
)

func fatorial(n uint) uint {
	switch {
	case n == 0:
		return 1
	default:
		return n * fatorial(n-1)
	}
}

// main is the entry point for the application.
func main() {
	resultado := fatorial(5)
	fmt.Println(resultado)

	//resutado2 := fatorial(-2)

}
