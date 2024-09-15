package data_structures

type linkedListNode struct {
	data int
	next *linkedListNode
}

type LinkedList struct {
	head   *linkedListNode
	lenght int
}

func (l *LinkedList) InsertAtHead(v int) {
	tmp := &linkedListNode{v, nil}

	if l.head == nil {
		l.head = tmp
	} else {
		tmp2 := l.head
		l.head = tmp
		tmp.next = tmp2
	}

	l.lenght++
}

func (l *LinkedList) InsertAtTail(v int) {
	tmp := &linkedListNode{v, nil}

	if l.head == nil {
		l.head = tmp
	} else {
		tmp2 := l.head
		for tmp2.next != nil {
			tmp2 = tmp2.next
		}
		tmp2.next = tmp
	}

	l.lenght++
}

func (l *LinkedList) Insert(idx, v int) {
	if idx == 0 {
		l.InsertAtHead(v)
	} else if idx == l.lenght-1 { //ZERO BASED INDEX DO NOT FORGET
		l.InsertAtTail(v)
	} else {
		tmp := &linkedListNode{v, nil}
		tmp2 := l.head

		for i := 0; i < idx-1; i++ {
			tmp2 = tmp.next
		}

		tmp.next = tmp2.next
		tmp2.next = tmp

		l.lenght++
	}
}

func (l *LinkedList) DeleteHead() {
	tmp := l.head
	l.head = tmp.next
	l.lenght--
}
func (l *LinkedList) DeleteTail() {
	tmp1 := l.head
	var tmp2 *linkedListNode
	for tmp1.next != nil {
		tmp2 = tmp1
		tmp1 = tmp1.next
	}
	tmp2.next = nil
	l.lenght--
}

func (l *LinkedList) DeleteAt(n int) {
	if n == 0 {
		l.DeleteHead()
	} else if n == l.lenght-1 {
		l.DeleteTail()
	} else {
		tmp1 := l.head
		for i := 0; i < n-1; i++ {
			tmp1 = tmp1.next
		}
		tmp2 := tmp1.next
		tmp1.next = tmp2.next
		l.lenght--
	}
}

func (l *LinkedList) Reverse() {
	var curr, prev, next *linkedListNode
	curr = l.head
	prev = nil

	for curr != nil {
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next
	}

	l.head = prev
}
