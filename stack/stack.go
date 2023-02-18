package stack

import "fmt"

type Stack struct {
	items []string
}

func New() *Stack {
	return &Stack{items: nil}
}

func (s *Stack) Push(data string) {
	s.items = append(s.items, data)
}

func (s *Stack) Pop() string {
	if s.IsEmpty() {
		return ""
	}

	s.items = s.items[:len(s.items)-1]

	if len(s.items) == 0 {
		return ""
	}

	return s.items[len(s.items)-1]
}

func (s *Stack) Top() string {
	if s.IsEmpty() {
		return ""
	}
	return s.items[len(s.items)-1]
}

func (s *Stack) IsEmpty() bool {
	if len(s.items) == 0 {
		return true
	}
	return false
}

func (s *Stack) Count() int {
	return len(s.items)
}

func (s *Stack) Print() {
	for _, item := range s.items {
		fmt.Print(item, " ")
	}
	fmt.Println()
}
