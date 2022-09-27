package main

import "fmt"

type Infomatiable interface {
	Info()
}

type User struct {
	Name string
	ID   string
	Age  int
}

func (user *User) Info() {
	fmt.Printf("유저: %s ID: %s 나이: %d\n", user.Name, user.ID, user.Age)
}

type VIPUser struct {
	UserInfo User
	VIPLevel int
	Price    int
}

func (user *VIPUser) Info() {
	fmt.Printf("VIP 유저: %s ID: %s 나이: %d VIP 레벨: %d VIP 가격: %d만 원\n", user.UserInfo.Name, user.UserInfo.ID, user.UserInfo.Age, user.VIPLevel, user.Price)
}

func main() {
	var x Infomatiable

	x = &User{"송하나", "hana", 23}
	x.Info()
	x = &VIPUser{User{"화랑", "hwarang", 40}, 3, 250}
	x.Info()
	fmt.Printf("%b", &x)
}
