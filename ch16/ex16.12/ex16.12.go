package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	s := []int{5, 2, 6, 4, 7, 3, 9}
	k := []float64{34.2, 84.22, 12.1}
	sort.Ints(s)
	sort.Float64s(k)
	slices.Reverse(s)
	slices.Reverse(k)
	fmt.Println(s)
	fmt.Println(k)
}
