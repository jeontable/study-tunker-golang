package main

import "fmt"

func main() {
	var char rune = '한'
	var char2 int32 = '함'

	fmt.Printf("%T\n", char) //rune == int32(4바이트)
	fmt.Println(char)
	fmt.Printf("%c\n", char)
	fmt.Println()
	fmt.Printf("%T\n", char2)
	fmt.Println(char2)
	fmt.Printf("%c\n", char2)
}
