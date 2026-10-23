package main

import "fmt"

func main() {
	var slice = make([]int, 3, 5)
	fmt.Println(slice)
	var slice2 = append(slice, 4)
	fmt.Println(slice)
	fmt.Println(slice2)
}
