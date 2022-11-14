package stack

import (
	"errors"
)

type Stack struct {
	top    int
	length int
	data   []int
}

func Create(n int) *Stack {
	return &Stack{
		top:    0,
		length: n,
		data:   make([]int, n+1),
	}
}

func (s *Stack) Empty() bool {
	return s.top == 0
}

func (s *Stack) Push(x int) {
	if s.top > s.length-1 {
		panic(errors.New("overflow"))
	}

	s.top = s.top + 1
	s.data[s.top] = x
}

func (s *Stack) Pop() int {
	if s.Empty() {
		panic(errors.New("underflow"))
	}

	s.top = s.top - 1
	return s.data[s.top+1]
}
