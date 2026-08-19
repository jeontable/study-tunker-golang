package main

import "fmt"

func main() {
	var x1 int8 = 34
	var x2 int16 = 34
	var x3 uint8 = 34
	var x4 uint16 = 34

	var x5 uint8 = 1
	var x6 uint8 = 2

	fmt.Printf("%d = %5d,\t %08b\n", x1, ^x1, uint8(^x1))
	fmt.Printf("%d = %5d,\t %016b\n", x2, ^x2, uint16(^x2))
	fmt.Printf("%d = %5d,\t %08b\n", x3, ^x3, ^x3)
	fmt.Printf("%d = %5d,\t %016b\n", x4, ^x4, ^x4)

	fmt.Printf("%d&%d = %5d,\t %08b\n", x5, x6, x5&x6, x5&x6)
	fmt.Printf("%d|%d = %5d,\t %08b\n", x5, x6, x5|x6, x5|x6)
}
