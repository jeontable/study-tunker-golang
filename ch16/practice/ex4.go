package main

import (
	"fmt"
	"sort"
)

type Athlete struct {
	Name  string
	Age   int
	Score int
	Pass  int
}

type Athletes []Athlete

func (a Athletes) Len() int {
	return len(a)
}

func (a Athletes) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a Athletes) Less(i, j int) bool {
	return a[i].Score < a[j].Score
}

func main() {
	athletes := []Athlete{
		{Name: "나통키", Age: 13, Score: 45, Pass: 1},
		{Name: "오맹태", Age: 16, Score: 24, Pass: 1},
		{Name: "오동도", Age: 18, Score: 54, Pass: 0},
		{Name: "황금산", Age: 16, Score: 36, Pass: 0},
	}
	sort.Sort(sort.Reverse(Athletes(athletes)))
	fmt.Println(athletes)
}
