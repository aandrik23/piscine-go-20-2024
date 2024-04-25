package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushFront(l *List, data interface{}) {
	n := &NodeL{Data: data}
	if l.Head == nil && l.Tail == nil {
		l.Head = n
		l.Tail = n
	} else {
		l.Head.Next = n
		l.Head = n
	}
}
