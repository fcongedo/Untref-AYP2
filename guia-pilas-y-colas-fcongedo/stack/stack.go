// Package stack proporciona una implementación de una pila genérica sobre un arreglo dinámico.
// Expone la estructura y funciones básicas para manipular una pila.
package stack

import "errors"

// Stack proporciona una pila cuyos elementos son de un tipo genérico.
// La implementación se basa en un arreglo dinámico.
type Stack[T any] struct {
	data []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{data: []T{}}
}

func (s *Stack[T]) Push(x T) {
	s.data = append(s.data, x)
}

func (s *Stack[T]) Pop() (T, error) {

	var x T
	if s.IsEmpty() {
		return x, errors.New("pila vacía")
	}
	x = s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]

	return x, nil

}

func (s *Stack[T]) Top() (T, error) {
	var x T
	if s.IsEmpty() {
		return x, errors.New("pila vacía")
	}
	x = s.data[len(s.data)-1]
	return x, nil
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}
