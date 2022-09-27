package main

import "fmt"

func main() {
	var minimumWage int = 10
	var workingHour int = 20
	var income int = minimumWage * workingHour

	fmt.Println(fmt.Sprintf("MinimumWage : %d / WorkingHour : %d / Income : %d", minimumWage, workingHour, income))
	fmt.Println(&income)
}
