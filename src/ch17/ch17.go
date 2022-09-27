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

func generateRandomValue(size int) (a int) {
	return rand.Intn(size)
}

func validateResult(src, tar int) bool {
	var result bool = false
	switch {
	case src < tar:
		fmt.Println("입력하신 숫자가 더 작습니다.")
	case src > tar:
		fmt.Println("입력하신 숫자가 더 큽니다.")
	default:
		fmt.Print("숫자를 맞췄습니다.")
		result = true
	}
	return result
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var target int = generateRandomValue(50000)
	var input int
	for i := 1; ; i++ {
		input = readInputValue()
		if validateResult(input, target) {
			fmt.Printf("축하합니다. 시도한 횟수: %d", i)
			break
		}
	}
	return
}
