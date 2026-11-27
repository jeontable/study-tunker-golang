package main

import "fmt"

func main() {
	var i int = 0

	var f func() = func() {
		i += 10
	}

	i++

	fmt.Println(i)

	f()

	fmt.Println(i)
}
