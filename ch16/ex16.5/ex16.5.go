package main

import (
	"fmt"
	"unsafe"
)

func main() {
	slice1 := make([]int, 3, 5)
	slice2 := append(slice1, 4, 5)

	fmt.Println(len(slice1), cap(slice1))
	fmt.Println(len(slice2), cap(slice2))
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Printf("%p\n", unsafe.SliceData(slice1))
	fmt.Printf("%p\n", unsafe.SliceData(slice2))

	slice1[1] = 100

	fmt.Println(len(slice1), cap(slice1))
	fmt.Println(len(slice2), cap(slice2))
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Printf("%p\n", unsafe.SliceData(slice1))
	fmt.Printf("%p\n", unsafe.SliceData(slice2))
}
