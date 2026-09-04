package main

import "fmt"

func main() {
	str1 := "안녕하세요. 한글 문자열입니다."
	str2 := str1 //string struct 크기만큼 복제된다. sting.data 자체가 복제되지 않는다.

	fmt.Printf(str1)
	fmt.Printf("\n")
	fmt.Printf(str2)
}
