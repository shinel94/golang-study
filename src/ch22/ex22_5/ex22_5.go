package main

import (
	"fmt"
	"unsafe"
)

func main() {
	m := make(map[string]string)
	m["이화랑"] = "서울시 광진구"
	m["송하나"] = "서울시 강남구"
	m["백두산"] = "부산시 사하구"
	m["최번개"] = "전주시 덕진구"

	m["최번개"] = "청주시 상당구"

	fmt.Printf("송하나의 주소는 %s입니다.\n", m["송하나"])
	fmt.Printf("백두산의 주소는 %s입니다.\n", m["백두산"])

	for k, v := range m {
		fmt.Println(k, v)
	}
	fmt.Printf("Empty Value : %s\n", m["Not Exist Key"])
	if v, ok := m["TEST"]; ok {
		fmt.Println("OK", v)
	} else {
		fmt.Println("No")
	}

	fmt.Println("Size of Map : ", unsafe.Sizeof(m))
	m["TEST"] = "TEST"
	fmt.Println("Size of Map : ", unsafe.Sizeof(m))
}
