package main

import "fmt"

const (
	RED int8 = 1 << iota
	GREEN
	BLUE
)

func main() {
	var a int8 = 1
	var b int8 = 1
	var c int8 = 1

	println(RED)
	println(GREEN)
	println(BLUE)

	println("======")
	fmt.Printf("%d %08b\n", a<<0, a<<0)
	fmt.Printf("%d %08b\n", b<<1, b<<1)
	fmt.Printf("%d %08b\n", c<<2, c<<2)
}
