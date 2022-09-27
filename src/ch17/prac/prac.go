package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var stdin = bufio.NewReader(os.Stdin)

func readInputValue() (a int) {
	for {
		fmt.Print("숫자를 입력하세요:")
		_, err := fmt.Scanln(&a)
		if err != nil {
			fmt.Println("숫자가 아닙니다.")
			stdin.ReadString('\n')
		} else {
			break
		}
	}

	return
}

func machineGame() {
	var initMoney int = 1000
	for {
		src := readInputValue()
		tar := rand.Intn(5) + 1
		if src == tar {
			initMoney += 500
			fmt.Printf("당첨!! 잔액 : %d\n 결과 %d / %d \n", initMoney, src, tar)
		} else {
			initMoney -= 100
			fmt.Printf("실패!! 잔액 : %d\n 결과 %d / %d \n", initMoney, src, tar)
		}
		if initMoney > 5000 || initMoney < 0 {
			break
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	machineGame()
}
