package main

import "strings"

type myString string

func (m myString) ToLower() myString {
	str := strings.ToLower(string(m))
	return myString(str)
}

func (m myString) ToUpper() myString {
	str := strings.ToUpper(string(m))
	return myString(str)
}

func main() {
	msg := myString("hello Go World")

	msg2 := msg.ToLower()
	msg3 := msg.ToUpper()
	println(msg2)
	println(msg3)
}
