package main

import "fmt"

func PrintAvgScore(name string, math int, eng int, history int) {
	total := math + eng + history

	avg := total / 3
	fmt.Println(name, " 님 평균 점수는 ", avg, "입니다.")
}

func Divide(a, b int) (result int, success bool) {
	if b == 0 {
		result = 0
		success = false
	} else {
		result = a / b
		success = true
	}
	return
}

func printNo(n int) {
	if n <= 0 {
		return
	}
	fmt.Println(n)
	printNo(n - 1)
	fmt.Println("After", n)
}

func main() {
	PrintAvgScore("김일등", 80, 74, 95)
	PrintAvgScore("송이등", 88, 92, 53)
	PrintAvgScore("박삼등", 78, 73, 78)

	fmt.Println(Divide(10, 3))
	fmt.Println(Divide(2, 3))
	fmt.Println(Divide(10, 0))

	printNo(10)
}
