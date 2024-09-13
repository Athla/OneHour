# Slow and Steady wins the Race
***
- Today was a hard day at the job with many frustrating moments, making me very tired.
- That being said, didn't code a lot, more study and doing proof of concepts to develop the file ingest system from another application, while designing it in Excalidraw

- Anyway, reinforced the knowledge of more basic structures such as Queue, Stack and Tree, to set the BFS and DFS in Binary Trees, it will also help  in the future.
- While doing that, I noticed that I was generating many helper types to compliment different structures, so I'll default to int as values for the Node.
```go
type Node struct {
    Value *int
}
```
- Futhermore, what I did was implement the stack and it's methods:
```go
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
```
- By doing so, I've broken the previous tests, so I already have a objective for tomorrow. 

# Considerations

- I would be lying if I said that I'm not sad that I couldn't keep up the same momentum from yesterday, or that I'm not disappointed in my because of that.
- However, commiting to learn is also commiting to some failures and slips along the way, since progress is not a linear function.
- PS: I'm thinking of chainging my current distro, I've using kUbuntu for a while but Hyprland and other tiling managers got my attention. What to do?
