package main

func soma1(a, b int) int {
	a = 50
	b = 50
	return a + b
}

func soma2(a, b *int) int {
	*a = 50
	*b = 50
	return *a + *b
}

func main() {
	minhaVar1 := 10
	minhaVar2 := 20

	soma1(minhaVar1, minhaVar2)

	println(minhaVar1, minhaVar2)
	
	soma2(&minhaVar1, &minhaVar2)

	println(minhaVar1, minhaVar2)
}
