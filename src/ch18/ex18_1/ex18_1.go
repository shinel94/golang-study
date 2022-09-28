package main

import "fmt"

func main() {
	var slice []int

	if len(slice) == 0 {
		fmt.Println("slice is empty", slice)
	}

	// slice[0] = 10
	// fmt.Println(slice)

	var s1 = []int{1, 2, 3}
	var s2 = [...]int{1, 2, 3}

	fmt.Printf("%T %T", s1, s2)
	var s3 = make([]int, 5)
	s3[3] = 10
	fmt.Println(s3)

	s3 = append(s3, 1,2,3,4,5)
	fmt.Println(s3)
	s3 = append(s3, []int{1,2,3,4,}...)
	fmt.Println(s3)
}
