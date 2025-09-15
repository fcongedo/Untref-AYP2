// Package queue implementa una cola genérica sobre un arreglo dinámico.
// Expone la estructura Queue y sus métodos para manipular una cola.
package queue

import "errors"

// Queue implementa una cola genérica sobre un arreglo dinámico.
type Queue[T any] struct {
	data []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{data: []T{}}
}

func (q *Queue[T]) Enqueue(v T) {
	q.data = append(q.data, v)
}

func (q *Queue[T]) Dequeue() (T, error) {
	// Implementar
	var x T
	if q.IsEmpty() {
		return x, errors.New("cola vacía")
	}
	x = q.data[0]
	q.data = q.data[1:]
	return x, nil
}

func (q *Queue[T]) Front() (T, error) {
	var x T
	if q.IsEmpty() {
		return x, errors.New("cola vacía")
	}
	return q.data[0], nil
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.data) == 0
}
