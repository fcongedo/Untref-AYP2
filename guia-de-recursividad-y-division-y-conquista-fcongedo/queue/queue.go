package queue

type Queue[T any] struct {
	elements []T
}

func New[T any](elements ...T) *Queue[T] {
	if elements == nil {
		elements = make([]T, 0)
	}
	return &Queue[T]{elements: elements}
}

func (q *Queue[T]) Enqueue(element T) {
	q.elements = append(q.elements, element)
}

func (q *Queue[T]) Dequeue() T {
	if len(q.elements) == 0 {
		panic("Queue is empty")
	}
	element := q.elements[0]
	q.elements = q.elements[1:]
	return element
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.elements) == 0
}

func (q *Queue[T]) Values() []T {
	return q.elements
}
