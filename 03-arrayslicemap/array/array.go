package main

import "fmt"

func main() {

	var notas [3]float64
	notas[0], notas[1], notas[2] = 7.8, 4.3, 9.1

	total := 0.0
	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}

	media := total / float64(len(notas))
	fmt.Printf("Media: %.2f\n", media)

	// Array
	var names [3]string
	names[0] = "Muhammad"
	names[1] = "Yusuf"
	names[2] = "Ardian"

	// Print Array
	println(names[0])
	println(names[1])
	println(names[2])
}
