package list

import (
	"testing"
)

func TestListShouldBeEmptyWhenCreated(t *testing.T) {
	l := New()

	if l.Size() == 0 && l.IsEmpty() && l.Head() == nil {
		return
	}
	t.Error("List size should be 0")
}

func TestWhenAnElementIsPrependedItShouldBeTheHead(t *testing.T) {
	l := New()
	previousSize := l.Size()

	l.Prepend("Foo").Prepend("Bar")

	if (l.Size() == previousSize+2) && l.Head() == "Bar" {
		return
	}
	t.Errorf("The head of the list should be Bar but it was %s and the size was %d instead of 2", l.Head(), l.Size())
}

func TestWhenAnElementIsDeletedShouldnBePresent(t *testing.T) {
	l := New()
	previousSize := l.Size()

	l.Append("1").Append("2").Append("3").Delete(1)
	if l.Size() == previousSize+2 {
		return
	}
	t.Errorf("The size should be 2 instead of %d", l.Size())
}

func TestWhenAnElementIsAppendedItShouldBeTheLastElement(t *testing.T) {
	l := New()
	previousSize := l.Size()

	l.Append("Foo")
	l.Append("Bar")
	if (l.Size() == previousSize+2) && l.Head() == "Foo" {
		return
	}
	t.Errorf("The head of the list should be Foo but it was %s and the size was %d instead of 2", l.Head(), l.Size())
}

func TestWhenGetNthElementItShouldBeCorrect(t *testing.T) {
	l := New()

	l.Prepend("1").Prepend("2").Prepend("3")
	value := l.Get(2)
	if value == "3" {
		return
	}
	t.Errorf("The element should be '3' but it was %s", value)
}
