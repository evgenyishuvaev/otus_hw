package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	cnt   int
	front *ListItem
	back  *ListItem
}

func (l *list) Len() int {
	return l.cnt
}

func (l *list) Front() *ListItem {
	return l.front
}

func (l *list) Back() *ListItem {
	return l.back
}

func (l *list) PushFront(value interface{}) *ListItem {
	l.cnt++

	newItem := ListItem{
		Value: value,
		Prev:  nil,
		Next:  nil,
	}

	curFront := l.front
	if curFront == nil {
		l.front = &newItem
		l.back = &newItem
		return &newItem
	}
	newItem.Next = l.front
	l.front.Prev = &newItem
	l.front = &newItem
	return &newItem
}

func (l *list) PushBack(value interface{}) *ListItem {
	l.cnt++

	newItem := ListItem{
		Value: value,
		Prev:  nil,
		Next:  nil,
	}

	if l.back == nil {
		l.front = &newItem
		l.back = &newItem
		return &newItem
	}
	newItem.Prev = l.back
	l.back.Next = &newItem
	l.back = &newItem
	return &newItem
}

func (l *list) MoveToFront(listItem *ListItem) {
	switch {
	case listItem == l.front:
		return
	case listItem == l.back:
		l.back.Prev.Next = nil
		l.back = l.back.Prev
	default:
		listItem.Next.Prev = listItem.Prev
		listItem.Prev.Next = listItem.Next
	}

	listItem.Prev = nil
	listItem.Next = l.front
	l.front.Prev = listItem
	l.front = listItem
}

func (l *list) Remove(listItem *ListItem) {
	switch {
	// single elem in list
	case l.front == listItem && l.Len() == 1:
		l.front = nil
		l.back = nil
	case listItem == l.front:
		listItem.Next.Prev = nil
		l.front = listItem.Next
	case listItem == l.back:
		listItem.Prev.Next = nil
		l.back = listItem.Prev
	// middle elem in list
	case listItem.Prev != nil && listItem.Next != nil:
		listItem.Prev.Next = listItem.Next
		listItem.Next.Prev = listItem.Prev
	}
	l.cnt--
}

func NewList() List {
	return new(list)
}
