package main

import "fmt"

type Carro struct {
	nome            string
	velocidadeAtual int
}

type Ferrari struct {
	Carro       // Campos anônimos
	turboLigado bool
}

func main() {
	f := Ferrari{}
	f.nome = "F40"
	f.velocidadeAtual = 0
	f.turboLigado = true

	// Acessando os campos do tipo anônimo
	fmt.Println(f.nome)
	fmt.Println(f.velocidadeAtual)
	fmt.Println(f.turboLigado)
}
