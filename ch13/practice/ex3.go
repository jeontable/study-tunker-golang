package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	FirstName string //string은 2개의 필드라서 16바이트
	LastName  string
	Age       int // int64 -> 8바이트
}

func main() {
	a := User{"길동", "홍", 23}
	fmt.Println(a)
	fmt.Println(unsafe.Sizeof(a))
}

//string 구조체는 data, len을 갖는다. data는 point변수(uintptr) 8바이트, len은 int이므로 8바이트
//어쨋든 string 변수의 구조체는 16바이트이다.
