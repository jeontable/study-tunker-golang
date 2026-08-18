package main

import "fmt"

func main() {
	var a = 324.13455
	var c = 3.14

	fmt.Printf("%08.2f\n", a)
	fmt.Printf("%8.2f\n", a)
	fmt.Printf("%8.2g\n", a)
	fmt.Printf("%8.5g\n", a)
	fmt.Printf("%08.5g\n", a)
	fmt.Printf("%f\n", a)      // default로 소숫점 이하 6개
	fmt.Printf("%-15.2f\n", a) // default로 소숫점 이하 6개
	fmt.Printf("%f\n", c)      // default로 소숫점 이하 6개
}
