package main

import (
	"fmt"
	"os"
)

func main() {

	f, err := os.Create("test.text")
	if err != nil {
		fmt.Println("Failed to create a file")
		return
	}

	defer func() {
		fmt.Println("반드시 호출됩니다.")
		f.Close()
		fmt.Println("파일을 닫았습니다")
	}()

	fmt.Println("파일에 Hello World를 씁니다.")
	fmt.Fprintln(f, "Hello World")
}
