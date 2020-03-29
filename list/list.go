package list

type item struct {
	value interface{}
	next *item
	previous *item
}

type list struct {
	head item
	size int64
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
func (l *list) Size() int64 {
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

// Adds an element to the end of the list
func (l *list) Prepend(s string) {
	var i = item{s, &l.head,&l.head}

	l.head.previous = &i
	l.head = i
	l.size += 1
}





