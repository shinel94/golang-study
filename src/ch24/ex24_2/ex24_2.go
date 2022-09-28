package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var wg sync.WaitGroup

func SumAtoB(a, b int) {
	sum := 0
	for i := a; i < b+1; i++ {
		sum += i
	}
	fmt.Printf("%d부터 %d까지 합계는 %d입니다.\n", a, b, sum)
	wg.Done()
}

func main() {
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go SumAtoB(1, rand.Intn(1500000000))
	}
	fmt.Println("Start Wait")
	wg.Wait()
	fmt.Println("Done Wait")
	fmt.Println("모든 계산이 완료되었습니다.")
}
