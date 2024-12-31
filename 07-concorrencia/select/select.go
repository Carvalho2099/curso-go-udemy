package main

import (
	"fmt"
	"time"
)

func oMaisRapido(url1, url2, url3 string) string {
	c1 := titulo(url1)
	c2 := titulo(url2)
	c3 := titulo(url3)

	// Estrutura de controle específica para concorrência
	// Select é um tipo de switch que seleciona o canal que está pronto para receber/enviar dados
	// O select não é bloqueante, ele tenta executar todos os cases
	// Se mais de um case estiver pronto, ele escolhe um aleatoriamente
	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(1000 * time.Millisecond):
		return "Todos perderam!"
	}
}

func main() {
	campeao := oMaisRapido(
		"https://www.google.com",
		"https://www.amazon.com",
		"https://www.youtube.com",
	)

	fmt.Println("Quem respondeu foi:", campeao)
}
