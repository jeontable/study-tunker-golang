package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, len(slice1))

	for i := range slice1 {
		slice2[i] = slice1[i]
	}

	fmt.Println(slice1)
	fmt.Println(slice2)

	slice3 := []int{1, 2, 3, 4, 5}
	slice4 := append([]int{}, slice3...)

	fmt.Println(slice4)
}
