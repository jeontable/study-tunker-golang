package main

import "fmt"

type opFunc func(int, int) int

func add(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

func getOperator(op string) opFunc {
	if op == "+" {
		return add
	} else if op == "*" {
		return mul
	} else {
		return nil
	}
}

func main() {
	var operator opFunc
	operator = getOperator("+")
	var a int = operator(3, 4)
	fmt.Println(a)
}
