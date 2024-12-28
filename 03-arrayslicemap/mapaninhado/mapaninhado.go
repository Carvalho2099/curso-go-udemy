package main

import "fmt"

func main() {
	// map de map
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriel": 1234.12,
			"Guga":    1234.12,
		},
		"J": {
			"José": 1234.12,
		},
		"P": {
			"Pedro": 1234.12,
		},
	}

	delete(funcsPorLetra, "P")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
		for nome, salario := range funcs {
			fmt.Println(nome, salario)
		}
	}
}
