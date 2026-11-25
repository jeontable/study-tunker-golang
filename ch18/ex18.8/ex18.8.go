package main

import "fmt"

type Stringer interface {
	String() string
}

type Student struct {
	Age string
}

func (s *Student) String() string {
	s.Age = s.Age + "========"
	return s.Age
}

func main() {
	var stringer Stringer
	stringer = &Student{"fdsafdsa"}
	a := stringer.(*Student)
	fmt.Println(stringer.String())
	fmt.Println(a.Age)
}
