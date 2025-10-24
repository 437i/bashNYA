package stack

import "fmt"

type Stack struct {
	items []int
}

func StackInit() *Stack {
	return &Stack{items: make([]int, 0)}
}

func (s *Stack) Push(num int) {
	s.items = append(s.items, num)
}

func (s *Stack) Pop() int {
	if s.IsEmpty() {
		fmt.Println("error: empty stack")
		return 0
	}
	last_index := len(s.items) - 1
	last_item := s.items[last_index]
	s.items = s.items[:last_index]
	return last_item
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Size() int {
	return len(s.items)
}

func (s *Stack) Clear() {
	s.items = s.items[:0]
}
