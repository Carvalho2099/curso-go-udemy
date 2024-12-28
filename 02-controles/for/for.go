package main

func main() {
	for i := 0; i <= 10; i++ {
		println(i)
	}

	i := 0
	for i <= 10 {
		println(i)
		i++
	}

	for i := 0; i <= 10; i += 2 {
		println(i)
	}

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			println(i, "par")
		} else {
			println(i, "ímpar")
		}
	}

	for {
		println("Para sempre...")
	}
}
