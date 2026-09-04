package main

import "fmt"

func main() {
	str1 := "가나다라마"
	str2 := "abcde"

	fmt.Printf("len(str1) = %d\n", len(str1)) // 문자의 길이가 아니라 찬지하는 바이트가 나온다.
	fmt.Printf("len(str2) = %d\n", len(str2)) // abcde는 5바이트
}
