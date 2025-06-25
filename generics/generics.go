package generics

import (
	"errors"
)

type Stack[T any] struct {
	values []T
}

var ErrPopOnEmpty = errors.New("stack is empty")

func (s *Stack[T]) Pop() (result T, error error) {
	if s.IsEmpty() {
		return result, ErrPopOnEmpty
	}

	result = s.values[len(s.values)-1]

	s.values = s.values[:len(s.values)-1]

	return
}

func (s *Stack[T]) Push(value T) {
	s.values = append(s.values, value)
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.values) == 0
}
