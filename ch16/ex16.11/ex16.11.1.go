package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}

	slice = append(slice, 0)
	copy(slice[3:], slice[2:])
	slice[2] = 100

	fmt.Println(slice)
}
