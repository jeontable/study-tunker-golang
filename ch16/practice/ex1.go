package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	slice := make([]int, 2, 3)
	copy(slice, array[1:3])
	slice = append(slice, 100)
	fmt.Println(array)
	fmt.Println(slice)
}
