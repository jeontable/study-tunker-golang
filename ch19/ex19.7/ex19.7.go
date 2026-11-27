package main

import (
	"fmt"
	"os"
)

type Writer func(string)

func writeHello(writer Writer) {
	writer("Hello World")
}

func main() {
	f, err := os.Create("test")
	if err != nil {
		fmt.Println("파일 생성 실패")
		return
	}

	defer f.Close()
	writeHello(func(msg string) {
		fmt.Fprintln(f, msg)
	})
}
