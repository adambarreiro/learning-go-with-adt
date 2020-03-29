// This package implements a linked list. It should be very similar to the implementation provided in
// the container package of the Go language, although it lacks of some functions and probably performance is
// worst.
//
// This was done just as a programming exercise to learn Go and gain some agility with its syntax.
//
package list

import (
	"fmt"
)

// ---- Data types -----------------------------------------------------------------------------------------------------

type item struct {
	value    interface{}
	next     *item
	previous *item
}

type list struct {
	head item
	size int
}

// ---- Public members -------------------------------------------------------------------------------------------------

// Creates a new empty list.
// O(1)
func New() *list {
	return new(list).init()
}

// Returns the size of the list.
// O(1)
func (l *list) Size() int {
	return l.size
}

// Returns true if the list has no elements.
// O(1)
func (l *list) IsEmpty() bool {
	return l.Size() == 0
}

// Returns the first element of the list
// O(1)
func (l *list) Head() interface{} {
	return l.head.value
}

// Returns the last element of the list
// O(1)
func (l *list) Butt() interface{} {
	return l.head.previous.value
}

// Adds an element to the beginning of the list
// O(1)
func (l *list) Prepend(element interface{}) *list {
	var newItem = item{element, &l.head, l.head.previous}

	l.head.previous.next = &newItem
	l.head.previous = &newItem
	l.head = newItem
	l.size++
	return l
}

// Adds an element at the end of the list
// O(1)
func (l *list) Append(element interface{}) *list {
	if l.IsEmpty() {
		l.Prepend(element)
		return l
	}
	var newItem = item{element, &l.head, l.head.previous}

	l.head.previous.next = &newItem
	l.head.previous = &newItem
	l.size++
	return l
}

// Obtains the element in the given position. If the position is out of bounds, returns nil
// O(n)
func (l *list) Get(position int) interface{} {
	if position < 0 || position >= l.Size() {
		return nil
	}

	i := 0
	var pointer = &l.head

	for i < position {
		pointer = l.head.next
		i++
	}
	return pointer.value
}

// Deletes an element at given position from the list. If the position is out of bounds, simply does nothing.
// O(n)
func (l *list) Delete(position int) *list {
	if position < 0 || position >= l.Size() {
		return l
	}

	i := 0
	var pointer = &l.head
	for i < position {
		pointer = l.head.next
		i++
	}
	pointer.previous.next = pointer.next
	pointer.next.previous = pointer.previous
	pointer.value = nil
	l.size--
	return l
}

// Prints the list in the standard output
// O(n)
func (l *list) Print() {
	fmt.Print(l.getPrintable())
}

// Prints the list in the standard output
// O(n)
func (l *list) Println() {
	fmt.Println(l.getPrintable())
}

// ---- Private members ------------------------------------------------------------------------------------------------

// Initializes the list when a new one is requested.
// O(1)
func (l *list) init() *list {
	l.size = 0
	l.head.next = &l.head
	l.head.previous = &l.head
	return l
}

// Gets a printable string from the list
// O(n)
func (l *list) getPrintable() string {
	output := "["
	if !l.IsEmpty() {
		pointer := &l.head
		i := 0

		for i < l.Size() {
			output += fmt.Sprintf("%v, ", pointer.value)
			pointer = pointer.next
			i++
		}
		output = fmt.Sprintf("%s", output[0:len(output)-2])
	}
	output += "]"
	return output
}
