package stack

import "testing"

func TestStackCount(t *testing.T) {
	stack := Stack{items: []string{"1", "2", "3"}}
	var count int
	count = stack.Count()
	if 3 != count {
		t.Error(
			"Test Count: ",
			"expected", 3,
			"got", count,
		)
	}
}

func TestStackPush(t *testing.T) {
	stack := Stack{items: []string{"1", "2", "3"}}
	var count int
	stack.Push("4")

	count = stack.Count()
	if 4 != count {
		t.Error(
			"Test Push: ",
			"expected", 4,
			"got", count,
		)
	}
}

func TestStackTop(t *testing.T) {
	stack := Stack{items: []string{"1", "2", "3"}}
	top := stack.Top()

	if "3" != top {
		t.Error(
			"Test Push: ",
			"expected", "3",
			"got", top,
		)
	}
}

func TestStackPop(t *testing.T) {
	stack := Stack{items: []string{"1", "2", "3"}}
	pop := stack.Pop()

	if "3" != pop {
		t.Error(
			"Test Pop: ",
			"expected", "3",
			"got", pop,
		)
	}

	count := stack.Count()
	if 2 != count {
		t.Error(
			"Test count after pop: ",
			"expected", 2,
			"got", count,
		)
	}
}

func TestStackIsEmpty(t *testing.T) {
	stack := Stack{items: []string{"1", "2", "3"}}

	if false != stack.IsEmpty() {
		t.Error(
			"Test IsEmpty: ",
			"expected", false,
			"got", true,
		)
	}
}
