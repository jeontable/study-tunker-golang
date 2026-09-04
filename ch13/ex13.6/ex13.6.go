package main

import "fmt"

func main() {
	str := "Hello 월드"
	runes := []rune(str)

	fmt.Printf("len(str) = %d\n", len(str))     //문자의 바이트 수
	fmt.Printf("len(runes) = %d\n", len(runes)) //배열의 요소 개수

}
