package main

import "fmt"

type Stringer interface {
	String() string
}

type Student struct {
	Age int
}

func (s *Student) String() string {
	return fmt.Sprintf("Studnet Age:%d\n", s.Age)
}

type Another struct{}

func (a *Another) String() string {
	return "Another Struct"
}

func PrintAge(stringer Stringer) {

	defer func() {
		v := recover()
		fmt.Println("recovered:", v)
	}()

	s := stringer.(*Student)
	fmt.Printf("Age: %d\n", s.Age)
}

func main() {
	s := &Student{15}
	fmt.Println(s.String())
	PrintAge(s)
	PrintAge(&Another{})
}
