package main

import "fmt"

type House struct {
	Address string
	Size    int
	Price   float64
	Type    string
}

func (house *House) Info() {
	fmt.Println("주소:", house.Address)
	fmt.Printf("크기: %d평\n", house.Size)
	fmt.Printf("가격: %.2f억 원\n", house.Price)
	fmt.Println("타입:", house.Type)
}

func main() {
	var house House = House{
		Address: "서울시 강동구...",
		Size:    28,
		Price:   9.8,
		Type:    "아파트",
	}
	house.Info()
}
