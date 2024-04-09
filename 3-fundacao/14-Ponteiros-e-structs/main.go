package main

import "fmt"

type Cliente struct {
	nome string
}

func (c Cliente) andou1() {
	c.nome = "Wesley Williams"
	fmt.Printf("O cliente %v andou1 \n", c.nome)
}

func (c *Cliente) andou2() {
	c.nome = "Wesley Williams"
	fmt.Printf("O cliente %v andou2 \n", c.nome)
}

type Conta struct {
	saldo int
}

func NewConta() *Conta {
	return &Conta{saldo: 0}
}

func (c Conta) simular1(valor int) int {
	c.saldo += valor
	println("Saldo1 = ", c.saldo)
	return c.saldo
}

func (c *Conta) simular2(valor int) int {
	c.saldo += valor
	println("Saldo2 = ", c.saldo)
	return c.saldo
}

func main() {
	wesley := Cliente{
		nome: "Wesley",
	}
	wesley.andou1()
	fmt.Printf("O valor da struct com nome %v \n", wesley.nome)
	wesley.andou2()
	fmt.Printf("O valor da struct com nome %v \n", wesley.nome)

	conta := Conta{saldo: 100}

	conta.simular1(200)
	println("Saldo1 = ", conta.saldo)

	conta.simular2(200)
	println("Saldo2 = ", conta.saldo)
}
