package main

type trackType map[string]track

func new_node[T trackType](element T) *Node[T] {
	return &Node[T]{
		element: element,
		next:    nil,
	}
}

func NewLinkedList[T trackType]() LinkeListI[T] {
	return &LinkedList[T]{
		head:  nil,
		count: 0,
	}
}

func (lst *LinkedList[T]) push(element T) {
	var current *Node[T]
	node := new_node(element)

	if lst.head == nil {
		lst.head = node
	} else {
		current = lst.head
		for current.next != nil {
			current = current.next
		}
		current.next = node
	}

	lst.count++
}
