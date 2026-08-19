package main

import "fmt"

func main() {
	var a float32 = 0.1
	var b float32 = 0.2
	var c float32 = 0.3

	fmt.Printf("%f + %f == %f : %v\n", a, b, c, a+b == c)
	fmt.Println(a + b)
}
