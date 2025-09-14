package stack

type Stack[T comparable] struct {
	// Implementar
}

func NewStack[T comparable]() *Stack[T] {
	// Implementar
	return nil
}

func (s *Stack[T]) Push(dato T) {
	// Implementar
}

func (s *Stack[T]) Pop() (T, error) {
	// Implementar
	var x T
	return x, nil
}

func (s *Stack[T]) Top() (T, error) {
	// Implementar
	var x T
	return x, nil
}

func (s *Stack[T]) IsEmpty() bool {
	// Implementar
	return false
}
