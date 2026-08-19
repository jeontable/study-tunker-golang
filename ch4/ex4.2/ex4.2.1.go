package main

import "fmt"

func main() {
	var a int8 = 5
	var b int8 = 3

	fmt.Printf("%8b\n", a^b)

	var c uint8 = 255
	var d uint8 = 16
	fmt.Printf("%08b\n", c)
	fmt.Printf("%08b\n", d)
	fmt.Printf("%08b\n", c&^d)
	fmt.Printf("%08b %d\n", c<<2, c<<2)
	fmt.Printf("%08b %d\n", c>>2, c>>2)
}
