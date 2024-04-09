package main

type MyNumber int

type Number interface {
	~int | ~float64
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func Compara[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	m1 := map[string]int{"Wesley": 1000, "João": 2000, "Maria": 3000}
	m2 := map[string]float64{"Wesley": 100.20, "João": 2000.3, "Maria": 300.0}

	m3 := map[string]MyNumber{"Wesley": 1000, "João": 2000, "Maria": 3000}
	println("Soma(int)\t= ", Soma(m1))
	println("Soma(float64)\t= ", Soma(m2))
	println("Soma(MyNumber)\t= ", Soma(m3))
	println("Compara(10, 10)\t= ", Compara(10, 10))
}
