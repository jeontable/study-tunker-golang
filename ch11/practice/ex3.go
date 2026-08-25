package main

import (
	"fmt"
	"unsafe"
)

type Padding struct {
	C float64
	E int
	B int32
	F float32
	D uint16
	A int8
	G int8
}

func main() {
	var p Padding
	fmt.Println(unsafe.Sizeof(p))
}
