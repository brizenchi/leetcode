package store

type Stack struct {
	head *Node
	size int
}
type Node struct {
	value int
	next  *Node
}

func (s *Stack) isEmpty() bool {
	return s.size == 0
}
func (s *Stack) Pop() (int, bool) {
	if s.isEmpty() {
		return 0, false
	}
	v := s.head.value
	s.head = s.head.next
	s.size--
	return v, true
}
func (s *Stack) Push(v int) {
	s.head = &Node{
		value: v,
		next:  s.head,
	}
	s.size++
}

type StackArr []int

func (s *StackArr) isEmpty() bool {
	return len(*s) == 0
}

func (s *StackArr) pop() (int, bool) {
	if s.isEmpty() {
		return 0, false
	}
	v := (*s)[0]
	*s = (*s)[:len(*s)-1]
	return v, true
}

func (s *StackArr) push(v int) {
	*s = append(*s, v)
}
