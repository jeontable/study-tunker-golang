package main

import "fmt"

type Attacker interface {
	Attack()
}

func main() {
	var att Attacker
	fmt.Println(att)
	att.Attack()
}
