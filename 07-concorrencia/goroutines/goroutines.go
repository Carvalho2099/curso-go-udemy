package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
		time.Sleep(time.Second)
	}
}

func main() {
	//fale("Maria", "Por que você não fala comigo?", 3)
	//fale("João", "Só posso falar depois de você!", 1)

	// go fale("Maria", "Ei...", 500)
	// go fale("João", "Opa...", 500)

	go fale("Maria", "Entendi!!!", 10)
	fale("João", "Parabéns!", 5)
}
