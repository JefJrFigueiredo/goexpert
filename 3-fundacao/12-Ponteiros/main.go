package main

func main() {
	// Memória -> Endereço -> Valor

	a := 10
	println("\t a = ", a)
	println("\t &a = ", &a)
	
	var ponteiro *int = &a
	println("var ponteiro *int = &a")
	println("\t ponteiro = ", ponteiro)
	
	*ponteiro = 20
	println("*ponteiro = 20")
	println("\t a = ", a)
	
	b := &a
	println("b := &a")
	println("\t b = ", b)
	println("\t *b = ", *b)

	*b = 30
	println("*b = 30")
	println("\t a = ", a)
}
