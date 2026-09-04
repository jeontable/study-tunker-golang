package main

import (
	"fmt"
	"unsafe"
)

// func ChangeString(str3 string) {
// 	str3[4] = 'T'
// }

func main() {
	str := "Hello World"
	str2 := str // string에 대한 struct 정보가 복사되는 것이지. struct.Data가 복사되는 것이 아니다.
	var str3 []byte = []byte(str)

	fmt.Println(unsafe.StringData(str))
	fmt.Println(unsafe.StringData(str2))
	fmt.Println(str3)
	fmt.Printf("%s", str3)

	//ChangeString(str)
}
