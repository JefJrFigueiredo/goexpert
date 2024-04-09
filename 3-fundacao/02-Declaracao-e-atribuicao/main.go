package main

const a = "Hello, World!"

var (
	b bool
	c int
	d string
	e float64
)

func main() {
	println("a = ", a)
	a := "X" // string
	println("a = ", a)
		
	println("b = ", b)
	b = true
	println("b = ", b)
	
	println("c = ", c)
	c = 10
	println("c = ", c)
	
	println("d = ", d)
	d = "Wesley"
	println("d = ", d)
	
	println("e = ", e)
	e = 1.2
	println("e = ", e)

	x()
}

func x() {
	println()

	println("a = ", a)
	
	println("b = ", b)
	
	println("c = ", c)
	
	println("d = ", d)
	
	println("e = ", e)
}
