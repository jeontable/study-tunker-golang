package main

import "fmt"

type Reader interface {
	Read()
}

type Closer interface {
	Close()
}

type File struct {
}

func (f *File) Read() {
}

func ReadFile(reader Reader) {
	c, ok := reader.(Closer)
	fmt.Println(ok)
	if ok == false {
		fmt.Println("변환 불가")
		return
	}
	c.Close()
}

func main() {
	file := &File{}
	ReadFile(file)
}
