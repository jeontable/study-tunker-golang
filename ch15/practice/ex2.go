package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

var stdin = bufio.NewReader(os.Stdin)

func valueInput() (int, error) {
	var i int
	_, err := fmt.Scanln(&i)
	if err != nil {
		stdin.ReadString('\n')
	}
	return i, err

}

func main() {
	var money int = 1000
	fmt.Println("당신이 가진 돈은 1000원입니다. 1~5사이의 값을 입력하여 배팅하세요. 맞으면 1000원 틀리면 -500원입니다.")

	for true {
		var value = rand.Intn(5)

		fmt.Print("Please enter a value: ")
		n, err := valueInput()
		if err != nil {
			fmt.Println("숫자만 입력하시오!")
		} else {
			if n == value {
				fmt.Println("맞았습니다! 1000원을 획득했습니다.")
				money += 1000
			} else {
				fmt.Println("틀렸습니다! 500원을 잃었습니다.")
				money -= 500
			}
		}
		fmt.Printf("현재 잔액은 %d 원입니다.\n", money)
		if money <= 0 {
			fmt.Println("잔액이 0원이 되어 게임을 종료합니다.")
			break
		}
		if money >= 5000 {
			fmt.Println("잔액이 5000원이 되어 게임을 종료합니다.")
		}
	}
}
