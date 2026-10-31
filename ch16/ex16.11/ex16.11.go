package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}

	slice = append(slice[:2], append([]int{100}, slice[2:]...)...)
	fmt.Println(slice)
}
