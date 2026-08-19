package main

import "fmt"

func main() {
	var x int8 = 1
	fmt.Println(x)
	fmt.Printf("%08b\n", x)

	x <<= 7
	fmt.Println(x)
	fmt.Printf("%08b\n", x)

	x >>= 7
	fmt.Println(x)
	fmt.Printf("%08b\n", uint8(x))
}
