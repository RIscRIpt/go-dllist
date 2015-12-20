# dllist
Implementation of dynamic *d*oubly *l*inked *list* in [Go language](https://github.com/golang)

## Usage
This dllist can be used as **stack** (LIFO), **queue** (FIFO), or just a as general purpose doubly linked list.

### Stack
```
import "github.com/RIscRIpt/dllist"
// ...
type item int
func (i item) Less(j interface{}) bool { return i < j.(item) }
func (i item) Equals(j interface{}) bool { return i == j.(item) }
// ...
stack := dllist.New()
stack.PushFront(item(1))
stack.PushFront(item(2))
stack.PushFront(item(3))
stack.PopFront() // 3, nil
stack.PopFront() // 2, nil
stack.PopFront() // 1, nil
stack.PopFront() // nil, error("List is empty!")
```

### Queue
```
import "github.com/RIscRIpt/dllist"
// ...
type item int
func (i item) Less(j interface{}) bool { return i < j.(item) }
func (i item) Equals(j interface{}) bool { return i == j.(item) }
// ...
queue := dllist.New()
queue.PushBack(item(1))
queue.PushBack(item(2))
queue.PushBack(item(3))
queue.PopFront() // 1, nil
queue.PopFront() // 2, nil
queue.PopFront() // 3, nil
queue.PopFront() // nil, error("List is empty!")
```

# TODO
 * Improve README.md
 * Make `Less` and `Equals` funcs of `Node` interface not required to implement to use dllist (if possible?)
