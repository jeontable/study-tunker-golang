package main

import "fmt"

func main() {
	age := 22

	if age < 10 {
		fmt.Println("어린이입니다")
	} else if age >= 20 && age < 30 {
		fmt.Println("Best age of your life")
	} else {
		fmt.Println("You are beautiful")
	}
}
