package main

import "fmt"

func main() {
	str := "Hello 월드!"
	arr := []rune(str)

	for i := 0; i < len(arr); i++ {
		fmt.Printf("타입:%T, 값:%d, 문자값:%c\n", arr[i], arr[i], arr[i])
	}
}

// string을 rune으로 변환하면 새로운 메모리 공간에 복사해서 rune을 만들기 때문에 메모리 낭비
// 하지만 rune으로 변환하지 않고 string을 바로 rang를 통해 읽어들이면 rune 단위로 읽기 때문에 메모리를 추가로 사용할 필요가 없다는 사실.
