package piscine

func ListPushFront(l *List, data interface{}) {
	n := &NodeL{Data: data}
	if l.Head == nil && l.Tail == nil {
		l.Head = n
		l.Tail = n
	} else {
		n.Next = l.Head
		l.Head = n
	}
}
