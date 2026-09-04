package main

import "fmt"

func main() {
	str := "Hello 월드"
	str2 := []rune(str)

	fmt.Printf("이 문자열의 타입: %T\n", str)
	fmt.Println(len(str))
	fmt.Println(len(str2))

	for i := 0; i < len(str); i++ {
		fmt.Printf("타입:%T, %d, 문자값:%c\n", str[i], str[i], str[i])
	}

	fmt.Println()

	for i := 0; i < len(str2); i++ {
		fmt.Printf("타입:%T, %d, 문자값:%c\n", str2[i], str2[i], str2[i])
	}
}
