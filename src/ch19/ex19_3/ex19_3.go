package main

import "fmt"

type account struct {
	balance   int
	firstName string
	lastName  string
}

func (a *account) withdrawPointer(amount int) {
	a.balance -= amount
}

func (a account) withdrawValue(amount int) {
	a.balance -= amount
}

func (a account) withdrawReturnValue(amount int) account {
	a.balance -= amount
	return a
}

func main() {
	var mainA *account = &account{100, "Joe", "Park"}
	mainA.withdrawPointer(30)
	fmt.Println(mainA.balance)

	mainA.withdrawValue(20) // (*mainA).withdrawValue(20)
	fmt.Println(mainA.balance)

	var mainB account = mainA.withdrawReturnValue(20)
	fmt.Println(mainB.balance)

	mainB.withdrawPointer(30) // (&mainB).withdrawPointer(30)
	fmt.Println(mainB)

}
