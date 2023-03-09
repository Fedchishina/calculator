package stack

import (
	"fmt"
	"log"
	"strconv"
)

type Stack struct {
	items []string
}

func (s *Stack) Push(data string) {
	s.items = append(s.items, data)
}

func (s *Stack) Pop() string {
	if s.IsEmpty() {
		return ""
	}

	val := s.items[len(s.items)-1]

	s.items = s.items[:len(s.items)-1]

	return val
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

func (s *Stack) PopNumber() float64 {
	var err error
	var number float64
	if s.Count() > 0 {
		number, err = strconv.ParseFloat(s.Pop(), 64)
		if err != nil {
			log.Fatalf("error during parsing %v to float : %v", number, err)
		}
	}

	return number
}
