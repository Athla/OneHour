package data_structures

type Stack struct {
	Values []*Node
}

func (s *Stack) Push(v *Node) {
	s.Values = append(s.Values, v)
}
func (s *Stack) Pop() *Node {
	if s.IsEmpty() {
		return nil
	}

	p := s.Values[len(s.Values)-1]
	s.Values = s.Values[:len(s.Values)-1]
	return p
}

func (s *Stack) IsEmpty() bool {
	if len(s.Values) == 0 {
		return false
	}
	return true
}
