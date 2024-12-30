package main

import "fmt"

type Item struct {
	produtoID int
	qtd       int
	preco     float64
}

type Pedido struct {
	userID int
	itens  []Item
}

func (p Pedido) valorTotal() float64 {
	total := 0.0
	for _, item := range p.itens {
		total += item.preco * float64(item.qtd)
	}
	return total
}

func main() {
	pedido := Pedido{
		userID: 1,
		itens: []Item{
			{1, 2, 12.10},
			{2, 1, 23.49},
			{11, 100, 3.13},
		},
	}
	fmt.Println(pedido.valorTotal())
}
