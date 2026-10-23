package main

import "fmt"

func main() {
	var slice1 = []int{1, 2, 3}
	slice2 := []int{1, 5: 2, 10: 3}
	slice3 := make([]int, 3)
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
}
