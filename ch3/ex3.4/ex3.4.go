package main

import "fmt"

func main() {
	str := "Hello\tGo\tWorld\n\"Go\"is Awesome!\n"

	fmt.Print(str)
	fmt.Printf("%s", str)
	fmt.Printf("%q", str)
}
