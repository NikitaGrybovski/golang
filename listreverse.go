package piscine_go

func ListReverse(l *List) {
	first := l.Head
	current := l.Head
	prev := l.Head
	prev = nil

	if l.Head == nil {
		return
	}

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	l.Head = prev
	l.Tail = first
	l.Tail.Next = nil
}
