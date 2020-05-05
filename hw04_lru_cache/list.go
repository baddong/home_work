package hw04_lru_cache //nolint:golint,stylecheck

type List interface {
	Len() int                      // длина списка
	Front() *Item                  // первый Item
	Back() *Item                   // последний Item
	PushFront(v interface{}) *Item // добавить значение в начало
	PushBack(v interface{}) *Item  // добавить значение в конец
	Remove(i *Item)                // удалить элемент
	MoveToFront(i *Item)           // переместить элемент в начало
}

type Item struct {
	Value interface{}
	Next  *Item
	Prev  *Item
}

type list struct {
	Head   *Item
	Tail   *Item
	Length int
}

func NewList() List {
	return &list{}
}

func (l *list) Len() int {
	return l.Length
}

func (l *list) Front() *Item {
	return l.Head
}

func (l *list) Back() *Item {
	return l.Tail
}

func (l *list) PushFront(v interface{}) *Item {
	newItem := &Item{Value: v}
	//check if list are empty and adds new item as head and tail
	if l.Head == nil {
		l.Head = newItem
		l.Tail = newItem
	} else {
		// point old back item to a new head item
		l.Head.Next = newItem
		newItem.Prev = l.Head
		l.Head = newItem
		l.Head.Next = nil
	}
	l.Length++
	return l.Head
}

func (l *list) PushBack(v interface{}) *Item {
	newItem := &Item{Value: v}
	//check if list are empty and adds new item as head and tail
	if l.Head == nil {
		l.Head = newItem
		l.Tail = newItem
	} else {
		// point old back item to a new back item
		l.Tail.Prev = newItem
		newItem.Next = l.Tail
		l.Tail = newItem
		l.Tail.Prev = nil
	}

	l.Length++
	return l.Tail
}

func (l *list) Remove(i *Item) {
	// checks for it's a head, back or in-between item
	switch i {
	case l.Head:
		// if head item has prev item
		if l.Head.Prev != nil {
			i.Prev.Next = nil
			l.Head = i.Prev
		} else {
			// if only one item exists
			l.Head, l.Tail = nil, nil
		}

	case l.Tail:
		// if tail item has next item
		if l.Tail.Next != nil {
			l.Tail = i.Next
			i.Next.Prev = nil
		} else {
			// if only one item exists
			l.Head, l.Tail = nil, nil
		}

	// if in-between item, pointer next item of prev item points to the next item
	// and pointer of the prev item of the next node points to the prev item
	default:
		i.Prev.Next = i.Next
		i.Next.Prev = i.Prev
	}
	l.Length--
}

func (l *list) MoveToFront(i *Item) {
	// checks for it's a head, back or in-between item
	switch i {
	case l.Head:
		return
	case l.Tail:
		i.Next.Prev = nil
		l.Tail = i.Next
	default:
		i.Prev.Next = i.Next
		i.Next.Prev = i.Prev
	}
	l.Head.Next = i
	i.Prev = l.Head
	l.Head = i
	l.Head.Next = nil
}
