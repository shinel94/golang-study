package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)

	var number int
	for {
		fmt.Println("입력하세요.")
		_, err := fmt.Scanln(&number)
		if err != nil {
			fmt.Println("숫자를 입력하세요.")
			stdin.ReadString('\n')
			continue
		}
		fmt.Printf("입력하신 숫자는 %d 입니다.\n", number)

		if number%2 == 0 {
			break
		}
	}
	fmt.Println("for문이 종료됐습니다.")
	for i := 0; i < 5; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Print("\n")
	}

	dan := 2
	b := 1
	for {
		for {
			fmt.Printf("%d * %d = %d\n", dan, b, dan*b)
			b++
			if b == 10 {
				break
			}
		}
		b = 1
		dan++
		if dan == 10 {
			break
		}
	}
}
