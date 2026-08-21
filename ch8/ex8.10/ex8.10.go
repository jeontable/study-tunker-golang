package main

func main() {
	a := 2

	switch a {
	case 1:
		println("1입니다")
	case 2:
		println("2입니다")
		fallthrough
	case 3:
		println("3입니다")
	case 4:
		println("4입니다")
	default:
		println("1, 2, 3이 아닙니다")
	}
}
