package main

import "fmt"

func main() {
	var a int8 = 30

	a <<= 2
	//a += 8
	fmt.Println(a)
	fmt.Printf("%08b\n", a)
}
