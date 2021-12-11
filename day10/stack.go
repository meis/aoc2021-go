package main

type stack []rune

func (s *stack) Top() rune {
	return (*s)[len((*s))-1]
}

func (s *stack) Pop() (stack, rune) {
	top := s.Top()
	return (*s)[:len((*s))-1], top
}

func (s *stack) Push(el rune) stack {
	return append(*s, el)
}

func (s *stack) Empty() bool {
	return len((*s)) == 0
}
