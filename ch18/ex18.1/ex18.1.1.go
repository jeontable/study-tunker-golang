package main

import (
	"fmt"
)

type Stringer interface {
	String() string
}

type Student struct {
	Name string
	Age  int
}

func (s Student) String() string {
	return fmt.Sprintf("안녕 나는 %s라고 해. 나는 %d살이야", s.Name, s.Age)
}

func main() {
	var student Student
	student.Name = "홍길동"
	student.Age = 42
	//student := Student{"홍길동", 42}

	var stringer Stringer

	stringer = student

	fmt.Println(stringer.String())
}
