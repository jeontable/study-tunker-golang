package main

import "fmt"

func Print(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println("int 입니다.")
		case bool:
			fmt.Println("bool 입니다.")
		case float64:
			fmt.Println("float64 입니다")
		default:
			fmt.Println("기타 입니다.")
		}
	}
}

func main() {
	Print(23, false, 34.3)
}
