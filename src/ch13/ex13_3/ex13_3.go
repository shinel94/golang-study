package main

import (
	"fmt"
	"unsafe"
)

type Student struct {
	Age   int
	No    int
	Score float64
}

func PrintStudnet(s Student) {
	fmt.Printf("나이: %d 번호: %d 점수: %.2f\n", s.Age, s.No, s.Score)
}

func main() {
	var studnet = &Student{15, 23, 88.2}

	studnet2 := *studnet

	PrintStudnet(studnet2)

	fmt.Printf("%X", studnet)
	fmt.Printf("%b", studnet)
	fmt.Printf("%X", studnet2)
	fmt.Printf("%b", studnet2)
	fmt.Println(unsafe.Pointer(studnet))
}
