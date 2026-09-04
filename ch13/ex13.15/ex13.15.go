package main

import "fmt"

func main() {
	var str string = "Hello World"
	var slice []byte = []byte(str) // byte는 uint8

	slice[2] = 'K'

	fmt.Println(str)
	fmt.Printf("%s\n", slice)
}
