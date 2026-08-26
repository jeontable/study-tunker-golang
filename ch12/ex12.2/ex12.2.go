package main

import "fmt"

func main() {
	var a int = 10
	var b int = 20

	var p1 *int = &a
	var p2 *int = &b
	var p3 *int = &b
	var p4 *int

	fmt.Printf("p1 == p2 : %v\n", p1 == p2)
	fmt.Printf("p2 == p3 : %v\n", p2 == p3)

	fmt.Printf("%d\n", *p1)
	fmt.Printf("%p\n", p1)

	if p4 == nil {
		fmt.Println("p4 is nil")
	}
}
