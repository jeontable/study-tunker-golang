package main

import "fmt"

type Data struct {
	value int
	data  [200]int
}

func ChangeData(arg *Data) {
	arg.value = 999
	arg.data[100] = 999
}

func main() {
	var data Data
	ChangeData(&data) //메모리 주소 전달

	fmt.Printf("value = %d\n", data.value)
	fmt.Printf("data[100] = %d\n", data.data[100])

	a1 := new(Data)
	fmt.Println((*a1).value)
	fmt.Println(a1.data[100])
	fmt.Printf("a1의 메모리 주소: %p\n", a1)

	var a2 *Data = &Data{}
	fmt.Println((*a2).value)
	fmt.Println(a2.data[100])
	fmt.Printf("a2의 메모리 주소: %p\n", a2)
}
