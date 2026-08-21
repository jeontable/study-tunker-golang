package main

import "fmt"

func main() {
	a := 3

	switch a {
	case 1:
		fmt.Println("1입니다")
	case 2:
		fmt.Println("2입니다")
	case 3:
		fmt.Println("3입니다")
	default:
		fmt.Println("1, 2, 3이 아닙니다")
	}
}
