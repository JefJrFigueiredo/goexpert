package main

import (
	"fmt"
)

func main() {
	var minhaVar interface{} = "Wesley Willians"
	println(minhaVar.(string))
	result, ok := minhaVar.(int)
	fmt.Printf("O valor de result é %v e o resultado de ok é %v \n", result, ok)
	result2 := minhaVar.(int)
	fmt.Printf("O valor de result2 é %v", result2)
}
