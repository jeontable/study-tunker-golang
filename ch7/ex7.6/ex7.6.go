package main

import "fmt"

func getMyAge() (int, bool) {
	return 20, true
}

func main() {

	if age, ok := getMyAge(); ok && age < 20 {
		fmt.Println("청소년입니다", age)
	} else if ok && age < 30 {
		fmt.Println("청년입니다", age)
	} else if ok {
		fmt.Println("중년입니다", age)
	} else {
		fmt.Println("나이를 알 수 없습니다")
	}

	//fmt.Println("Your age is", age)
}
