package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type ferrarri struct {
	modelo          string
	turboLigado     bool
	velocidadeAtual int
}

func (f *ferrarri) ligarTurbo() {
	f.turboLigado = true
}

func main() {
	carro1 := ferrarri{"F40", false, 0}
	carro1.ligarTurbo()

	var carro2 esportivo = &ferrarri{"F40", false, 0}
	carro2.ligarTurbo()

	fmt.Println(carro1, carro2)
}
