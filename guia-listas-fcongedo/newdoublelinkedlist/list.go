package newdoublelinkedlist

// List es un lista enlazada simple con nodo ficticio
type List[T comparable] struct {
	cabeza  *Node[T]
	cola    *Node[T]
	tamanio int
}

// NewList crea una nueva lista enlazada
func NewList[T comparable]() *List[T] {
	// Implementar
	return nil
}

// Append agrega un nuevo nodo al final de la lista
func (l *List[T]) Append(dato T) {
	// Implementar
}

// Prepend agrega un nuevo nodo al principio de la lista
func (l *List[T]) Prepend(dato T) {
	// Implementar
}

// RemoveFirst elimina el primer nodo de la lista
func (l *List[T]) RemoveFirst() {
	// Implementar
}

// RemoveLast elimina el último nodo de la lista
func (l *List[T]) RemoveLast() {
	// Implementar
}

// Remove elimina el nodo con el dato pasado por parámetro
func (l *List[T]) Remove(dato T) {
	// Implementar
}

// IsEmpty devuelve true si la lista está vacía
func (l *List[T]) IsEmpty() bool {
	// Implementar
	return false
}

// Size devuelve la cantidad de nodos en la lista
func (l *List[T]) Size() int {
	// Implementar
	return -1
}

// Head devuelve el dato del primer nodo de la lista
func (l *List[T]) Head() *Node[T] {
	// Implementar
	return nil
}

// Tail devuelve el dato del último nodo de la lista
func (l *List[T]) Tail() *Node[T] {
	// Implementar
	return nil
}

// Find busca un nodo con el dato pasado por parámetro y lo devuelve
func (l *List[T]) Find(dato T) *Node[T] {
	// Implementar
	return nil
}

// Clear elimina todos los nodos de la lista
func (l *List[T]) Clear() {
	// Implementar
}
