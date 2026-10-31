package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name string
	Age  int
}

type Students []Student

func (s Students) Len() int {
	return len(s)
}

func (s Students) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}

func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	s := []Student{
		{"화랑", 31},
		{"선덕", 28},
		{"진흥", 35},
		{"무열", 40},
		{"경순", 25},
		{"태종", 50},
		{"세종", 45},
	}

	sort.Sort(Students(s))
	fmt.Println(s)
}
