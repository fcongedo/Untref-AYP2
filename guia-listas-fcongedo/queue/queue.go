package queue

type Queue[T comparable] struct {
	// Implementar
}

func NewQueue[T comparable]() *Queue[T] {
	// Implementar
	return nil
}

func (q *Queue[T]) Enqueue(dato T) {
	// Implementar
}

func (q *Queue[T]) Dequeue() (T, error) {
	// Implementar
	var x T
	return x, nil

}

func (q *Queue[T]) Front() (T, error) {
	// Implementar
	var x T
	return x, nil
}

func (q *Queue[T]) IsEmpty() bool {
	// Implementar
	return false
}
