package main

import "fmt"

func main() {
	var age = 22

	if age >= 10 && age <= 15 {
		fmt.Println("U R Young")
	} else if age > 30 || age < 20 {
		fmt.Println("U R not 20s")
	} else {
		fmt.Println("Best age of your life")
	}
}
