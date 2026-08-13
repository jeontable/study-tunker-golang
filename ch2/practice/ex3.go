package main

import "fmt"

func main() {
	var f1 float32 = 123.546789 * 345.678
	var f2 float32 = float32(123.546789) * 345.678
	fmt.Println(f1)
	fmt.Println(f2)
}