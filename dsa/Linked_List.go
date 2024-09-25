package dsa

type linkedListNode struct {
	data int
	next *linkedListNode
}

type LinkedList struct {
	Head   *linkedListNode
	Lenght int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		Lenght: 0,
	}
}

func (l *LinkedList) InsertAtHead(v int) {
	tmp := &linkedListNode{v, nil}

	if l.Head == nil {
		l.Head = tmp
	} else {
		tmp2 := l.Head
		l.Head = tmp
		tmp.next = tmp2
	}

	l.Lenght++
}

func (l *LinkedList) InsertAtTail(v int) {
	tmp := &linkedListNode{v, nil}

	if l.Head == nil {
		l.Head = tmp
	} else {
		tmp2 := l.Head
		for tmp2.next != nil {
			tmp2 = tmp2.next
		}
		tmp2.next = tmp
	}

	l.Lenght++
}

func (l *LinkedList) Insert(idx, v int) {
	if idx == 0 {
		l.InsertAtHead(v)
	} else if idx == l.Lenght-1 { //ZERO BASED INDEX DO NOT FORGET
		l.InsertAtTail(v)
	} else {
		tmp := &linkedListNode{v, nil}
		tmp2 := l.Head

		for i := 0; i < idx-1; i++ {
			tmp2 = tmp.next
		}

		tmp.next = tmp2.next
		tmp2.next = tmp

		l.Lenght++
	}
}

func (l *LinkedList) DeleteHead() {
	tmp := l.Head
	l.Head = tmp.next
	l.Lenght--
}
func (l *LinkedList) DeleteTail() {
	tmp1 := l.Head
	var tmp2 *linkedListNode
	for tmp1.next != nil {
		tmp2 = tmp1
		tmp1 = tmp1.next
	}
	tmp2.next = nil
	l.Lenght--
}

func (l *LinkedList) DeleteAt(n int) {
	if n == 0 {
		l.DeleteHead()
	} else if n == l.Lenght-1 {
		l.DeleteTail()
	} else {
		tmp1 := l.Head
		for i := 0; i < n-1; i++ {
			tmp1 = tmp1.next
		}
		tmp2 := tmp1.next
		tmp1.next = tmp2.next
		l.Lenght--
	}
}

func (l *LinkedList) Reverse() {
	var curr, prev, next *linkedListNode
	curr = l.Head
	prev = nil

	for curr != nil {
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next
	}

	l.Head = prev
}
