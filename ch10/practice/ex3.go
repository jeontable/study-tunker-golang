package main

import "fmt"

func ChangeArrar(arr [5]int) {
	arr[3] = 3000
}

func main() {
	a := [5]int{1, 2, 3, 4, 5}

	ChangeArrar(a)

	fmt.Println(a[3])
}
