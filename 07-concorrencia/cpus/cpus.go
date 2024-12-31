package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Número de CPUs:", runtime.NumCPU())
}
