package stack

type Stack struct {
	stack []int
	pos   int
}

func New() *Stack {
	return &Stack{}
}

func (s *Stack) Push(value int) {
	s.pos++
	s.stack = append(s.stack, value)
}

func (s *Stack) Pop() int {
	s.pos--
	res := s.stack[s.pos]
	s.stack = s.stack[:s.pos]
	return res
}
