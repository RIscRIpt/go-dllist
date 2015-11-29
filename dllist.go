package dllist

import (
	"errors"
)

type List struct {
	ptr    *Node
	length uint
}

func New() *List {
	return &List{}
}

func (list *List) initialize(node *Node) {
	node.next = node
	node.prev = node
	list.ptr = node
}

func (list *List) IsEmpty() bool {
	return list.ptr == nil
}

func (list *List) Length() uint {
	return list.length
}

func (list *List) First() *Node {
	return list.ptr
}

func (list *List) Last() *Node {
	return list.ptr.prev
}

func (list *List) NodeIsFirst(node *Node) bool {
	return list.First() == node
}

func (list *List) AssertNotEmpty() error {
	if list.IsEmpty() {
		return errors.New("List is empty!")
	} else {
		return nil
	}
}

func (list *List) RotateToNext() (err error) {
	if err = list.AssertNotEmpty(); err == nil {
		list.ptr = list.ptr.next
	}
	return
}

func (list *List) RotateToPrev() (err error) {
	if err = list.AssertNotEmpty(); err == nil {
		list.ptr = list.ptr.prev
	}
	return
}

func (list *List) PushFront(data Noder) {
	list.PushBack(data)
	list.RotateToPrev()
}

func (list *List) PushBack(data Noder) {
	new_node := &Node{
		data: data,
	}
	if !list.IsEmpty() {
		list.First().Prepend(new_node)
	} else {
		list.initialize(new_node)
	}
	list.length++
}

func (list *List) removeFirst() {
	first := list.First()
	if first.IsClosed() {
		list.ptr = nil
	} else {
		list.ptr = list.ptr.next
	}
	first.Break()
}

func (list *List) Remove(node *Node) {
	if !list.NodeIsFirst(node) {
		node.Break()
	} else {
		list.removeFirst()
	}
	list.length--
}

func (list *List) PopFront() (data Noder, err error) {
	if err = list.AssertNotEmpty(); err == nil {
		node := list.First()
		list.Remove(node)
		data = node.data
	}
	return
}

func (list *List) PopBack() (data Noder, err error) {
	if err = list.AssertNotEmpty(); err == nil {
		node := list.Last()
		list.Remove(node)
		data = node.data
	}
	return
}

func (list *List) FindAll(data Noder) (result []*Node) {
	if !list.IsEmpty() {
		for i, node := NewIterator(list.First(), list.Last()); node != nil; node = i.Next() {
			if node.data.Equals(data) {
				result = append(result, node)
			}
		}
	}
	return
}

func Sort(iter *Iterator) (new_first *Node) {
	pivot := iter.Curr()
	pivdat := pivot.data
	inserter := pivot
	iterator := iter.Next()
	for {
		for iterator != nil && !iterator.data.Less(pivdat) {
			iterator = iter.Next()
		}
		if iterator != nil {
			tmp_node := iterator
			iterator = iter.Next()
			if iterator == nil {
				iter.SetLast(tmp_node.prev)
			}
			inserter.Prepend(tmp_node)
			inserter = tmp_node //==inserter.prev
		} else {
			break
		}
	}
	new_first = inserter
	new_last := iter.Prev()
	if new_first != pivot && new_first.next != pivot {
		left_iter, _ := NewIterator(new_first, pivot.prev)
		new_first = Sort(left_iter)
	}
	if pivot != new_last && pivot.next != new_last {
		right_iter, _ := NewIterator(pivot.next, new_last)
		Sort(right_iter)
	}
	return
}

func (list *List) Sort() {
	if !list.IsEmpty() {
		iter, _ := NewIterator(list.First(), list.Last())
		list.ptr = Sort(iter)
	}
}
