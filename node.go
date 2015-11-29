package dllist

type Node struct {
	prev, next *Node
	data       Noder
}

type Noder interface {
	Less(other interface{}) bool
	Equals(other interface{}) bool
}

func (node *Node) Data() Noder {
	return node.data
}

func (node *Node) Prev() *Node {
	return node.prev
}

func (node *Node) Next() *Node {
	return node.next
}

func (node *Node) HasPrev() bool {
	return node.prev != nil
}

func (node *Node) HasNext() bool {
	return node.next != nil
}

func (node *Node) IsClosed() bool {
	return node.prev == node && node.next == node
}

func (node *Node) IsIsolated() bool {
	return !node.HasPrev() && !node.HasNext()
}

func (node *Node) Wrest() {
	if node.HasNext() {
		node.next.prev = node.prev
	}
	if node.HasPrev() {
		node.prev.next = node.next
	}
}

func (node *Node) Break() {
	node.Wrest()
	node.prev = nil
	node.next = nil
}

func (node *Node) Reunite() {
	if node.HasNext() {
		node.next.prev = node
	}
	if node.HasPrev() {
		node.prev.next = node
	}
}

func (where *Node) Prepend(what *Node) {
	what.Wrest()
	what.next = where
	what.prev = where.prev
	what.Reunite()
}

func (where *Node) Append(what *Node) {
	what.Wrest()
	what.next = where.next
	what.prev = where
	what.Reunite()
}
