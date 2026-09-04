package main

import (
	"fmt"
	"unsafe"
)

func main() {
	str1 := "Hello world"
	str2 := str1

	fmt.Println(str1)
	fmt.Println(str2)
	fmt.Printf("%s\n", str1)
	fmt.Printf("%s\n", str2)
	fmt.Printf("%p\n", &str1)
	fmt.Printf("%p\n", &str2)
	fmt.Println(unsafe.StringData(str1))
	fmt.Println(unsafe.StringData(str2))
}
