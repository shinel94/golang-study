package main

import (
	"container/list"
	"fmt"
)

type Queue struct {
	v *list.List
}

func (q *Queue) Push(value interface{}) {
	q.v.PushBack(value)
}
func (q *Queue) Pop() interface{} {
	front := q.v.Front()
	if front != nil {
		return q.v.Remove(front)
	}
	return nil
}

func GetQueue() *Queue {
	return &Queue{list.New()}
}

type Stack struct {
	v *list.List
}

func (s *Stack) Push(value interface{}) {
	s.v.PushBack(value)
}

func (s *Stack) Pop() interface{} {
	back := s.v.Back()
	if back != nil {
		return s.v.Remove(back)
	}
	return nil
}

func GetStack() *Stack {
	return &Stack{
		list.New(),
	}
}

func main() {
	queue := GetQueue()

	for i := 1; i < 5; i++ {
		queue.Push(i)
	}
	v := queue.Pop()
	for v != nil {
		fmt.Printf("%v -> ", v)
		v = queue.Pop()
	}

	stack := GetStack()

	for i := 1; i < 5; i++ {
		stack.Push(i)
	}
	v = stack.Pop()
	for v != nil {
		fmt.Printf("%v -> ", v)
		v = stack.Pop()
	}
}
