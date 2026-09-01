package main

import "fmt"

func main() {
	str := "Hello 월드"

	fmt.Printf("이 문자열의 타입: %T\n", str)

	for i := 0; i < len(str); i++ {
		fmt.Printf("타입:%T, 문자값:%c\n", str[i], str[i])
	}
}
