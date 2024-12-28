package main

import "time"

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "int"
	case float32, float64:
		return "float"
	case func():
		return "função"
	case bool:
		return "bool"
	case string:
		return "string"
	default:
		return "desconhecido"
	}
}

func main() {
	println(tipo(2))
	println(tipo(2.3))
	println(tipo(func() {}))
	println(tipo("oi"))
	println(tipo(true))
	println(tipo(time.Now()))
}
