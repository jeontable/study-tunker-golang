package main

import "fmt"

func main() {
	str := "Hello World"

	runes := []rune{72, 101, 108, 108, 111, 32, 87, 111, 114, 108, 100} //초기화는 중괄호로

	fmt.Println(str)
	fmt.Println(string(runes))

	fmt.Println([]rune(runes)) //slicing 한 것과 같은 효과

	for _, c := range []rune(runes) {
		fmt.Printf("%c", c)
	}
}
