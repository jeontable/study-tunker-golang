package main

import "fmt"

func main() {
	str1 := "BBB"
	str2 := "aaaaAAA"
	str3 := "BBAD"
	str4 := "ZZZ"
	str5 := "a"
	str6 := "A"

	fmt.Printf("%s > %s : %v\n", str1, str2, str1 > str2)   // false
	fmt.Printf("%s < %s : %v\n", str1, str3, str1 < str3)   // false
	fmt.Printf("%s <= %s : %v\n", str1, str4, str1 <= str4) // true
	fmt.Printf("%s <= %s : %v\n", str5, str6, str5 <= str6) // false
	fmt.Printf("%d\n", []rune(str5)[0])
	fmt.Printf("%d\n", []rune(str6)[0])
}

// 대문자가 앞에, 소문자는 뒤에 있다. 따라서 대문자 < 소문자이다.
