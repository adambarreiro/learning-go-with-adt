package list

import (
	"testing"
)

func TestListShouldBeEmptyWhenCreated(t *testing.T) {
	l := New()

	if l.Size() != 0 && l.IsEmpty() {
		t.Error("List size should be 0")
	}
}

func TestListShouldHaveOneMoreElementWhenElementIsAdded(t *testing.T) {
	l := New()
	previousSize := l.Size()

	l.Prepend("Foo")
	if l.Size() != previousSize + 1 {
		t.Errorf("List size should be 1 but it was %d", l.Size())
	}
}

func TestListHeadShouldBeTheElementWhenAnElementIsPrepended(t *testing.T) {
	l := New()

	l.Prepend("Foo")
	if (l.Head() != "Foo") {
		t.Errorf("The head of the list should be Foo but it was %s", l.Head())
	}
}