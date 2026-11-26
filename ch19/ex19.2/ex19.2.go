package main

import (
	"fmt"
	"os"
)

func main() {

	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("파일 생성이 실패했습니다.")
		return
	}

	defer fmt.Println("파일을 닫았습니다.")
	defer f.Close()
	defer fmt.Println("이제 파일을 닫을까 합니다.")

	fmt.Println("파일에 헬로우 월드를 씁니다")
	fmt.Fprintln(f, "Hello World")
}
