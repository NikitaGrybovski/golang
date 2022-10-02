package piscine_go

func ListLast(l *List) interface{} {
	if l.Head == nil {
		return nil
	}

	last := l.Head

	for last.Next != nil {
		last = last.Next
	}
	return last.Data
}
