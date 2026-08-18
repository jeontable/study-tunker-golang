package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)
	var a int
	var b int

	n, err := fmt.Scanln(&a, &b)
	println("n:", n, "err:", err)
	if err != nil {
		fmt.Println(n, err)
		stdin.ReadString('\n')
	} else {
		fmt.Println(n, a, b)
	}
}
