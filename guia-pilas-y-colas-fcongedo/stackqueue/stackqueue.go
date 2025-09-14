package stackqueue

// Implementación de una cola genérica utilizando dos pilas
type StackQueue[T any] struct {
	// Implementar
}

// Crea una nueva cola vacía.
func NewStackQueue[T any]() *StackQueue[T] {
	// Implementar
	return nil
}

// Agrega un elemento a la cola.
func (q *StackQueue[T]) Enqueue(v T) {
	// Implementar
}

// Elimina y devuelve el elemento al frente de la cola.
func (q *StackQueue[T]) Dequeue() (T, error) {
	// Implementar
	var x T
	return x, nil
}

// Devuelve el elemento al frente de la cola.
func (q *StackQueue[T]) Front() (T, error) {
	// Implementar
	var x T
	return x, nil
}

// Devuelve true si la cola está vacía.
func (q *StackQueue[T]) IsEmpty() bool {
	// Implementar
	return false
}
