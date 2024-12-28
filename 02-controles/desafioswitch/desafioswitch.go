package main

func notaParaCoonceito(n float64) string {
	switch {
	case n >= 9 && n <= 10:
		return "A"
	case n >= 7 && n < 9:
		return "B"
	case n >= 5 && n < 7:
		return "C"
	case n >= 3 && n < 5:
		return "D"
	default:
		return "E"
	}
}

func main() {
	println(notaParaCoonceito(9.8))
	println(notaParaCoonceito(6.9))
	println(notaParaCoonceito(2.1))
	println(notaParaCoonceito(11.0))
}
