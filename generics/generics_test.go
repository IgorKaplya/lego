package generics

import (
	"errors"
	"fmt"
	"testing"
)

func TestAssertFunctions(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		AssertEqual(t, 1, 1)
		AssertNotEqual(t, 1, 2)
	})
	t.Run("string", func(t *testing.T) {
		AssertEqual(t, "1", "1")
		AssertNotEqual(t, "1", "2")
	})
	t.Run("bool", func(t *testing.T) {
		AssertTrue(t, fmt.Sprint(1) == "1")
		AssertFalse(t, "1" == "2")
	})
}

func AssertEqual[T comparable](t testing.TB, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func AssertNotEqual[T comparable](t testing.TB, got, want T) {
	t.Helper()
	if got == want {
		t.Errorf("got %+v, didn't want %+v", got, want)
	}
}

func AssertTrue(t testing.TB, got bool) {
	t.Helper()
	AssertEqual(t, got, true)
}

func AssertFalse(t testing.TB, got bool) {
	t.Helper()
	AssertEqual(t, got, false)
}

func AssertError(t testing.TB, got error, want error) {
	t.Helper()
	if !errors.Is(got, want) {
		t.Errorf(`got "%v", want "%v"`, got, want)
	}
}

func TestStack(t *testing.T) {
	t.Run("int-stack", func(t *testing.T) {
		var stack = new(Stack[int])
		AssertTrue(t, stack.IsEmpty())

		stack.Push(123)
		AssertFalse(t, stack.IsEmpty())

		stack.Push(456)
		var poppedFirst, firstPopError = stack.Pop()
		AssertError(t, firstPopError, nil)
		AssertEqual(t, 456, poppedFirst)

		var poppedSecond, secondPopError = stack.Pop()
		AssertError(t, secondPopError, nil)
		AssertEqual(t, 123, poppedSecond)

		AssertTrue(t, stack.IsEmpty())
	})

	t.Run("string-stack", func(t *testing.T) {
		var stack = new(Stack[string])
		AssertTrue(t, stack.IsEmpty())

		stack.Push("123")
		AssertFalse(t, stack.IsEmpty())

		stack.Push("456")
		var poppedFirst, firstPopError = stack.Pop()
		AssertError(t, firstPopError, nil)
		AssertEqual(t, "456", poppedFirst)

		var poppedSecond, secondPopError = stack.Pop()
		AssertError(t, secondPopError, nil)
		AssertEqual(t, "123", poppedSecond)

		AssertTrue(t, stack.IsEmpty())
	})

	t.Run("pop-errors-when-empty", func(t *testing.T) {
		var stack = new(Stack[struct{}])

		var popped, err = stack.Pop()
		AssertError(t, err, ErrPopOnEmpty)
		AssertEqual(t, struct{}{}, popped)
	})
}
