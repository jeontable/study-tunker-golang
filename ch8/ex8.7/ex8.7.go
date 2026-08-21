package main

import "fmt"

func getMyAge() int {
	return 10
}

func main() {
	switch age := getMyAge(); {
	case age < 10:
		fmt.Println("어린이입니다", age)
	case age < 20:
		fmt.Println("청년입니다", age)
	case age < 30:
		fmt.Println("중년입니다", age)
	default:
		fmt.Println("나이를 알 수 없습니다", age)
	}
}
