package main

import "fmt"

func sum(nums ...int) int {
	sum := 0

	fmt.Printf("이 데이터의 타입은 %T 입니다\n", nums)
	for _, v := range nums {
		sum += v
	}

	return sum
}

func main() {
	fmt.Println(sum(1, 2, 3, 4))
	fmt.Println(sum(10, 20))
	fmt.Println(sum())
}
