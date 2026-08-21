package main

import "fmt"

func getMyAge() int {
	return 10
}

func main() {
	switch age := getMyAge(); age {
	case 10:
		fmt.Println("어린이입니다", age)
	case 33:
		fmt.Println("청년입니다", age)
	default:
		fmt.Println("나이를 알 수 없습니다", age)
	}
}
