package dllist

type Iterator struct {
	start, current, end *Node
}

func NewIterator(start, end *Node) (*Iterator, *Node) {
	return &Iterator{
		start:   start,
		current: start,
		end:     end,
	}, start
}

func (iter *Iterator) First() (node *Node) {
	return iter.start
}

func (iter *Iterator) SetFirst(node *Node) {
	iter.start = node
}

func (iter *Iterator) Last() (node *Node) {
	return iter.end
}

func (iter *Iterator) SetLast(node *Node) {
	iter.end = node
}

func (iter *Iterator) Next() (node *Node) {
	if iter.current == nil || iter.current == iter.end {
		node = nil
	} else {
		node = iter.current.next
	}
	iter.current = node
	return
}

func (iter *Iterator) Curr() (node *Node) {
	return iter.current
}

func (iter *Iterator) Prev() (node *Node) {
	if iter.current == nil {
		node = iter.end
	} else if iter.current != iter.start {
		node = iter.current.prev
	} else {
		node = nil
	}
	iter.current = node
	return
}
