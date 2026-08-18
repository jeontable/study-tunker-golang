package main

import "fmt"

func main() {
	var a int
	var b int

	fmt.Print("Please input two integers: ")
	n, err := fmt.Scan(&a, &b)
	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println(n, a, b)
	}
}
