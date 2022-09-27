package main

import "fmt"

type test struct {
	a int
	b int16
	c int32
	d int64
}

func main() {
	var a int = 3
	var b int
	var c = 4
	d := 5
	t := test{}
	fmt.Println(a, b, c, d, t)
}
