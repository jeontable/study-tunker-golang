package main

import "fmt"

func main() {
	str := "Hello 월드!" //int32 겠구먼
	for _, v := range str {
		fmt.Printf("타입:%T, 값:%d, 문자값:%c\n", v, v, v)
	}
}

//rune은 int32 - 4byte
//byte는 uint8 - 1byte
//range를 사용하여 string을 읽으면 rune 단위로 읽게 된다.
