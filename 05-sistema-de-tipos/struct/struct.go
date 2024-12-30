package main

import "fmt"

type Produto struct {
	nome     string
	preco    float64
	desconto float64
}

// Método: função com receiver (receptor)
func (p Produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {
	var produto1 Produto
	produto1 = Produto{
		nome:     "Lápis",
		preco:    1.79,
		desconto: 0.05,
	}
	// Acessando o método
	fmt.Println(produto1.precoComDesconto())

	produto2 := Produto{"Notebook", 1989.90, 0.10}
	fmt.Println(produto2.precoComDesconto())
}
