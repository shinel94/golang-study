package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5}

	slice2 := append([]int{}, slice1...)
	// slice2 := make([]int, len(slice1))

	// for i, v := range slice1 {
	// 	slice2[i] = v
	// }
	slice1[1] = 100

	fmt.Println(slice1)
	fmt.Println(slice2)

	slice1 = []int{1, 2, 3, 4, 5}
	slice2 = make([]int, 3, 10)
	slice3 := make([]int, 10)
	cnt1 := copy(slice2, slice1)
	cnt2 := copy(slice3, slice1)
	fmt.Println(cnt1, slice2)
	fmt.Println(cnt2, slice3)

	slice2 = make([]int, len(slice1))
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)
	fmt.Printf("%p %p", slice1, slice2)
}
