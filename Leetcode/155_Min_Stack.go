package leetcode

type MinStack struct {
	values []int
	min    []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (m *MinStack) Push(val int) {
	if len(m.min) == 0 || val <= m.GetMin() {
		m.min = append(m.min, val)
	}
	m.values = append(m.values, val)
}

func (m *MinStack) Pop() {
	if m.Top() == m.GetMin() {
		m.min = m.min[:len(m.min)-1]
	}

	m.values = m.values[:len(m.values)-1]
}

func (m *MinStack) Top() int {
	return m.values[len(m.values)-1]
}

func (m *MinStack) GetMin() int {
	return m.min[len(m.min)-1]
}
