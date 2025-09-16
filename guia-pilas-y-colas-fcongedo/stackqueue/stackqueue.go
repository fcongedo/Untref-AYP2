package stackqueue

import (
	"fmt"
	"untref-ayp2/guia-pilas-colas/stack"
)

// Implementación de una cola genérica utilizando dos pilas
type StackQueue[T any] struct {
	stackIn  *stack.Stack[T]
	stackOut *stack.Stack[T]
}

// Crea una nueva cola vacía.
func NewStackQueue[T any]() *StackQueue[T] {
	return &StackQueue[T]{
		stackIn:  stack.NewStack[T](),
		stackOut: stack.NewStack[T](),
	}
}

// Agrega un elemento a la cola.
func (q *StackQueue[T]) Enqueue(v T) {
	q.stackIn.Push(v)
}

// Elimina y devuelve el elemento al frente de la cola.
func (q *StackQueue[T]) Dequeue() (T, error) {
	var x T
	if q.stackOut.IsEmpty() {
		for !q.stackIn.IsEmpty() {
			v, _ := q.stackIn.Pop()
			q.stackOut.Push(v)
		}
	}
	if q.stackOut.IsEmpty() {
		return x, fmt.Errorf("cola vacía")
	}
	return q.stackOut.Pop()
}

// Devuelve el elemento al frente de la cola.
func (q *StackQueue[T]) Front() (T, error) {
	// Implementar
	var x T
	if q.stackOut.IsEmpty() {
		for !q.stackIn.IsEmpty() {
			v, _ := q.stackIn.Pop()
			q.stackOut.Push(v)
		}
	}
	if q.stackOut.IsEmpty() {
		return x, fmt.Errorf("cola vacía")
	}
	return q.stackOut.Top()

}

// Devuelve true si la cola está vacía.
func (q *StackQueue[T]) IsEmpty() bool {

	return q.stackIn.IsEmpty() && q.stackOut.IsEmpty()
}
