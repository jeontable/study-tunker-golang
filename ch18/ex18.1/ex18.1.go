package main

import "fmt"

type Stringer interface {
	String() string
}

type Student struct {
	Name string
	Age  int
}

func (s Student) String() string {
	return fmt.Sprintf("안녕! 나는 어린이 %d살 %s라고 해", s.Age, s.Name)
}

type Adult struct {
	Name string
	Age  int
}

func (a Adult) String() string {
	return fmt.Sprintf("안녕! 나는 어른 %d살 %s라고 해", a.Age, a.Name)
}

func main() {
	student := Student{Name: "철수", Age: 20}
	adult := Adult{Name: "영희", Age: 30}

	var stringer Stringer

	stringer = student

	fmt.Printf("%s\n", stringer.String())

	stringer = adult
	fmt.Printf("%s\n", stringer.String())
}
