package main

import "fmt"

func f() {
	fmt.Println("Start f()")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover panic in f() - ", r)
		}
	}()
	g()
	fmt.Println("Finish f()")
}

func g() {
	fmt.Printf("9 / 3 = %d\n", h(9, 3))
	fmt.Printf("9 / 0 = %d\n", h(9, 0))
}

func h(a, b int) int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover panic in h() - ", r)
			panic(r)
		}
	}()
	if b == 0 {
		panic("can't divide by zero")
	} else {
		return a / b
	}
}

func main() {
	f()
	fmt.Println("Done Main Function")
}
