package main

func main() {
	// Slice
	numeros := make([]int, 5)

	// Slice is a reference to an array
	println(numeros)

	numeros = make([]int, 10, 20)
	println(numeros, len(numeros), cap(numeros))
}
