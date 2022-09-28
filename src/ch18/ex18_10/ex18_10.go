package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	idx := 2

	for i := idx + 1; i < len(slice); i++ {
		slice[i-1] = slice[i]
	}
	slice = slice[:len(slice)-1]
	fmt.Println(slice)

	slice = append(slice[:idx], slice[idx+1:]...)
	fmt.Println(slice)

	slice = append(slice, 0)
	for i := len(slice) - 2; i >= idx; i-- {
		slice[i+1] = slice[i]
	}
	slice[idx] = 100
	fmt.Println(slice)
	slice = append(slice[:idx], append([]int{200}, slice[idx:]...)...)
	fmt.Println(slice)

	slice = append(slice, 0)
	copy(slice[idx+1:], slice[idx:])
	slice[idx] = 300
	fmt.Println(slice)
}
