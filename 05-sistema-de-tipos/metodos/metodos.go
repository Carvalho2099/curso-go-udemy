package main

import (
	"fmt"
	"strings"
)

type Pessoa struct {
	nome      string
	sobrenome string
}

func (p Pessoa) nomeCompleto() string {
	return p.nome + " " + p.sobrenome
}

func (p *Pessoa) setNomeCompleto(nomeComplero string) {
	partes := strings.Split(nomeComplero, " ")
	p.nome = partes[0]
	p.sobrenome = partes[1]
}

func main() {
	p1 := Pessoa{"João", "Pedro"}
	fmt.Println(p1.nomeCompleto())

	p1.setNomeCompleto("Maria Pereira")
	fmt.Println(p1.nomeCompleto())
}
