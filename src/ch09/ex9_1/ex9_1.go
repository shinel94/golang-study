package main

import "fmt"

var light string = "red"

func canGo() {
	if light == "green" {
		fmt.Println("길을 건넌다")
	} else {
		fmt.Println("대기한다")
	}
}

func airConditioning(temp int) {
	if temp > 28 {
		fmt.Println("에어컨을 켠다")
	} else if temp <= 3 {
		fmt.Println("히터를 켠다")
	} else {
		fmt.Println("대기한다.")
	}
}

func main() {
	canGo()
	light = "green"
	canGo()
	airConditioning(33)
	airConditioning(13)
	airConditioning(3)
}
