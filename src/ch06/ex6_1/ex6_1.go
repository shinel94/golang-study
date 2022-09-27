package main

import "fmt"

func main() {
	var x int32 = 8
	var y int32 = 3

	var s float32 = 3.14
	var t float32 = 5

	fmt.Println("x + y =", x+y)
	fmt.Println("x - y =", x-y)
	fmt.Println("x * y =", x*y)
	fmt.Println("x / y =", x/y)
	fmt.Println("x % y =", x%y)
	fmt.Println("x << y =", x<<y)
	fmt.Println("x >> y =", x>>y)

	fmt.Println("s * t = ", s*t)
	fmt.Println("s / t = ", s/t)
	fmt.Printf("%08b\n", 10)
	fmt.Printf("%08b\n", 10<<1)
	fmt.Printf("%08b\n", 10>>1)
	fmt.Printf("%08b\n", 10&^2)
	fmt.Printf("%08b\n", 10&^4)
	fmt.Printf("%08b\n", 10&^8)
	fmt.Printf("%08b\n", 100&^8)
	fmt.Printf("%08b\n", 100&^16)
	fmt.Printf("%08b\n", 100&^32)
	fmt.Printf("%08b\n", 100&^64)
}
