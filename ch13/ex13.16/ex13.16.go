package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var str string = "Hello World"
	var slice []byte = []byte(str)

	fmt.Printf("str:\t%p\n", unsafe.StringData(str))    // 문자열 struct 안에 있는 Data의 주소
	fmt.Printf("str struct:\t%p\n", &str)               // 문자열의 struct 포인터
	fmt.Printf("slice:\t%p\n", unsafe.SliceData(slice)) // slice struct 안에 있는 Data의 주소(?)
	fmt.Printf("slice struct:\t%p\n", &slice)           // slice struct 자체의 주소
}
