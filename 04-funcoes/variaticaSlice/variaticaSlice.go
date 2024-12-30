package main

func imprimaAprovados(aprovados ...string) {
	println("Lista de aprovados:")
	for i, aprovado := range aprovados {
		println(i+1, aprovado)
	}
}

func main() {
	aprovados := []string{"Maria", "Pedro", "Rafael", "Guilherme"}
	imprimaAprovados(aprovados...)
}
