package data_structures

type Stack struct {
	Values []*int
}

func (s *Stack) Push(v *int) {
	s.Values = append(s.Values, v)
}
func (s *Stack) Pop() *int {
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
