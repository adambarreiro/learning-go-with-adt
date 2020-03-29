package list

import (
	"testing"
)

func TestListShouldBeEmptyWhenCreated(t *testing.T) {
	l := New()

	if l.Size() == 0 && l.IsEmpty() && l.Head() == nil && l.Butt() == nil {
		return
	}
	t.Error("List size should be 0")
}

func TestHeadAndButtAreSameWhenEmptyList(t *testing.T) {
	l := New()

	if l.Head() == l.Butt() {
		return
	}
	t.Error("Head and butt don't match")
}

func TestHeadAndButtAreSameWhenOneElement(t *testing.T) {
	l := New()

	l.Append(0)
	if l.Head() == l.Butt() {
		return
	}
	t.Error("Head and butt don't match with the first element")
}

func TestHeadAndButtAreNotTheSameWhenMoreElements(t *testing.T) {
	l := New()

	l.Append(0).Append(1)
	if l.Head() != l.Butt() {
		return
	}
	t.Error("Head and butt match but the list has more than 1 element")
}

func TestPrependAndAppendAreSameWithEmptyList(t *testing.T) {
	l1, l2 := New(), New()

	l1.Append(0)
	l2.Prepend(0)
	if l1.Head() == l2.Head() && l1.Size() == l2.Size() {
		return
	}
	t.Error("Both lists with same elements should match")
}

func TestAppendAndDeleteShouldEqualToSameList(t *testing.T) {
	l := New()

	l.Append(0).Delete(0)
	if l.IsEmpty() && l.Head() == nil {
		return
	}
	t.Error("List should be empty when appending and deleting the element")
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
