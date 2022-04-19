package main

import "fmt"

type Stack struct {
	top  int
	size int
	data []int
}

func new(size int) (stack *Stack) {
	return &Stack{
		size: size,
		top:  -1,
		data: make([]int, size),
	}
}

func (s *Stack) push(item int) {
	if s.is_full() {
		return
	} else {
		s.top++
		s.data[s.top] = item
	}
}

func (s Stack) is_empty() bool {
	return s.top < 0
}

func (s Stack) is_full() bool {
	return s.top == s.size-1
}

func (s Stack) peek() (item int) {
	if s.is_empty() {
		return
	}

	return s.data[s.top]

}

func main() {

	stack := new(10)

	for i := 1; i <= 12; i++ {
		stack.push(i)
	}

	fmt.Println(stack.peek())

}
