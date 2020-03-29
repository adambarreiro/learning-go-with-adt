package list

import (
	"fmt"
)

type item struct {
	value interface{}
	next *item
	previous *item
}

type list struct {
	head item
	size int
}

// Creates a new empty list.
func New() *list {
	return new(list).init()
}

// Initializes the list when a new one is requested.
func (l *list) init() *list {
	l.size = 0
	l.head.next = &l.head
	l.head.previous = &l.head
	return l
}

// Returns the size of the list.
func (l *list) Size() int {
	return l.size
}

// Returns true if the list has no elements.
func (l *list) IsEmpty() bool {
	return l.size == 0
}

// Returns the first element of the list
func (l *list) Head() interface{} {
	return l.head.value
}

// Adds an element to the end of the list
func (l *list) Prepend(element interface{}) {
	var i = item{element, &l.head,&l.head}

	l.head.previous = &i
	l.head = i
	l.size += 1
}

// Obtains the element in the given position
func (l *list) Get(position int) (interface{}, error) {
	if position >= l.size {
		return nil, fmt.Errorf("provided position %d is out of bounds of this list with size %d", position, l.size)
	}

	i := 0
	var pointer = &l.head

	for i < position {
		i++
		pointer = l.head.next
	}
	return pointer.value, nil
}






