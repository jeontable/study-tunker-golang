package main

import "fmt"

type ColorType int

const (
	Red ColorType = iota
	Blue
	Green
	Yellow
)

func colorToString(color ColorType) string {
	switch color {
	case Red:
		fmt.Printf("%d\n", Red)
		return "빨강"
	case Green:
		fmt.Printf("%d\n", Green)
		return "초록"
	case Blue:
		fmt.Printf("%d\n", Blue)
		return "파랑"
	case Yellow:
		fmt.Printf("%d\n", Yellow)
		return "노랑"
	default:
		return "알 수 없음"
	}
}

func GetMyFaviroteColor() ColorType {
	return Yellow
}

func main() {
	fmt.Println("My favorite is", colorToString(GetMyFaviroteColor()))
}
