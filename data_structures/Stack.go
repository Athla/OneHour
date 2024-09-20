package data_structures

type Stack struct {
	Values []*int
	Lenght uint
}

func NewStack() *Stack {
	return &Stack{Lenght: 0}
}

func (s *Stack) Push(v *int) {
	s.Values = append(s.Values, v)
	s.Lenght++
}
func (s *Stack) Pop() *int {
	if s.IsEmpty() {
		return nil
	}

	p := s.Values[len(s.Values)-1]
	s.Values = s.Values[:len(s.Values)-1]
	s.Lenght--
	return p
}

func (s *Stack) IsEmpty() bool {
	if s.Lenght == 0 {
		return false
	}
	return true
}
