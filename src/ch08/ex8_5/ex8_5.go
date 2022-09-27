package main

import "fmt"

const PI = 3.14

func main() {
	var a int = PI * 100 // -> var a int = 314
	var f float64 = PI * 100 // -> var f float64 = 314

	fmt.Println(a, f, PI) // -> fmt.Println(a, f, 3.14)
}