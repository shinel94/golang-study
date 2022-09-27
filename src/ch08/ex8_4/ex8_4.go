package main

import "fmt"

const (
	Red int = iota
	Blue
	Green
)
const (
	C1 uint = iota + 1
	C2
	C3
)
const (
	B1 uint = 1 << iota
	B2
	B3
)

func main() {
	fmt.Println(Red, Blue, Green)
	fmt.Println(C1, C2, C3)
	fmt.Println(B1, B2, B3)
}
