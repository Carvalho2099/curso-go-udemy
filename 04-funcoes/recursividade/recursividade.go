package main

import (
	"errors"
	"fmt"
)

func fatorial(n int) (int, error) {
	switch {
	case n < 0:
		return -1, errors.New("Número inválido")
	case n == 0:
		return 1, nil
	default:
		fatorialAnterior, _ := fatorial(n - 1)
		return n * fatorialAnterior, nil
	}
}

// main is the entry point for the application.
func main() {
	resultado, _ := fatorial(5)
	fmt.Println(resultado)

	_, err := fatorial(-4)
	if err != nil {
		fmt.Println(err)
	}
}
