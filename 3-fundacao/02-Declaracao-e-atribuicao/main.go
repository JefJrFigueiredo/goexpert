package main

const a = "Hello, World!"

var (
	b bool
	c int
	d string
	e float64
)

func main() {
	print("a = "); println(a)
	a := "X" // string
	print("a = "); println(a)
		
	print("b = "); println(b)
	b = true
	print("b = "); println(b)
	
	print("c = "); println(c)
	c = 10
	print("c = "); println(c)
	
	print("d = "); println(d)
	d = "Wesley"
	print("d = "); println(d)
	
	print("e = "); println(e)
	e = 1.2
	print("e = "); println(e)

	x()
}

func x() {
	println()

	print("a = "); println(a)
	
	print("b = "); println(b)
	
	print("c = "); println(c)
	
	print("d = "); println(d)
	
	print("e = "); println(e)
}
