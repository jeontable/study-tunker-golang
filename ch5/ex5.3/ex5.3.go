package main

import "fmt"

func PrintAvgScore(name string, math int, eng int, history int) {
	total := math + eng + history
	avg := total / 3
	fmt.Println(name, "의 평균 점수는", avg, "점 입니다.")
}

func main() {
	PrintAvgScore("홍길동", 80, 74, 95)
	PrintAvgScore("김철수", 88, 92, 53)
	PrintAvgScore("이영희", 78, 73, 78)
}
