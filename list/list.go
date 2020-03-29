package list

import (
	"fmt"
)

type item struct {
	value    interface{}
	next     *item
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

// Returns the size of the list.
func (l *list) Size() int {
	return l.size
}

// Returns true if the list has no elements.
func (l *list) IsEmpty() bool {
	return l.Size() == 0
}

// Returns the first element of the list
func (l *list) Head() interface{} {
	return l.head.value
}

// Adds an element to the beginning of the list
func (l *list) Prepend(element interface{}) {
	var newItem = item{element, &l.head, &l.head}

	l.head.previous = &newItem
	l.head = newItem
	l.size++
}

// Adds an element at the end of the list
func (l *list) Append(element interface{}) {
	if l.IsEmpty() {
		l.head = item{element, l.head.next, l.head.previous}
		l.size++
		return
	}

	i := 0
	var pointer = &l.head

	for i < l.size {
		pointer = pointer.next
		i++
	}
	var newItem = item{element, nil, pointer}
	pointer.next = &newItem
	pointer.next.next = pointer.next
	l.size++
}

// Obtains the element in the given position
func (l *list) Get(position int) (interface{}, error) {
	if position >= l.Size() {
		return nil, fmt.Errorf("provided position %d is out of bounds of this list with size %d", position, l.size)
	}

	i := 0
	var pointer = &l.head

	for i < position {
		pointer = l.head.next
		i++
	}
	return pointer.value, nil
}

// Initializes the list when a new one is requested.
func (l *list) init() *list {
	l.size = 0
	l.head.next = &l.head
	l.head.previous = &l.head
	return l
}

// Prints the list in the standard output
func (l *list) Print() {
	fmt.Print(l.getPrintable())
}

// Prints the list in the standard output
func (l *list) Println() {
	fmt.Println(l.getPrintable())
}

// Gets a printable string from the list
func (l *list) getPrintable() string {
	output := ""
	if !l.IsEmpty() {
		output += "["
		pointer := &l.head
		i := 0
		for i < l.Size() {
			output += fmt.Sprintf("%s, ", pointer.value)
			pointer = pointer.next
			i++
		}
		output = fmt.Sprintf("%s]", output[0:len(output)-2])
	}
	return output
}
