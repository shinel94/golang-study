package main

import "fmt"

func main() {

	var a *int = new(int)
	var b *int = new(int)

	n, err := fmt.Scan(a, b)
	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println(n, a, b, &a, &b, *a, *b) // int값의 포인터 가 a,b로 전달됨, &a, &b역시 포인터로 그대로 전달됨(call by reference), *a, *b pointer의 주소값에 기록된 값을 반환함
	}
}
