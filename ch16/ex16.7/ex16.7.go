package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	slice := array[1:2]

	fmt.Println(slice)
	fmt.Printf("%T\n", array)
	fmt.Printf("%T\n", slice)
	fmt.Println(len(array), cap(array))
	fmt.Println(len(slice), cap(slice))

	array[1] = 100
	fmt.Println("After modifying array[1]:")
	fmt.Println(array)
	fmt.Println(len(slice), cap(slice))

	slice = append(slice, 500)
	fmt.Println("After appending to slice:")
	fmt.Println(array)
	fmt.Println(slice)
	fmt.Println(len(slice), cap(slice))

	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice2 := slice1[1:3:4]

	fmt.Println(slice2)
	fmt.Println(len(slice2), cap(slice2))
}
