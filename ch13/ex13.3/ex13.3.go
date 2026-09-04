package main

import "fmt"

func main() {
	var char rune = '한'

	fmt.Printf("%T\n", char) //rune == int32(4바이트)
	fmt.Println(char)
	fmt.Printf("%c\n", char)
}
