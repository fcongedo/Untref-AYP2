package list

// CircularList implementa una lista enlazada circular genérica.
type CircularList[T comparable] struct {
	head *DoubleLinkedNode[T]
	size int
}

// NewCircularList crea una nueva lista enlazada circular.
//
// Uso:
//
//	list := NewCircularList[int]() // Crea una nueva lista enlazada circular.
func NewCircularList[T comparable]() *CircularList[T] {
	// Implementar
	return nil
}

// Head devuelve el primer nodo de la lista.
//
// Uso:
//
//	head := list.Head() // Obtiene el primer nodo de la lista.
//
// Retorna:
//   - el primer nodo de la lista.
func (l *CircularList[T]) Head() *DoubleLinkedNode[T] {
	// Implementar
	return nil
}

// Tail devuelve el último nodo de la lista.
//
// Uso:
//
//	tail := list.Tail() // Obtiene el último nodo de la lista.
//
// Retorna:
//   - el último nodo de la lista.
func (l *CircularList[T]) Tail() *DoubleLinkedNode[T] {
	// Implementar
	return nil
}

// Size devuelve el tamaño de la lista.
//
// Uso:
//
//	size := list.Size() // Obtiene el tamaño de la lista.
//
// Retorna:
//   - el tamaño de la lista.
func (l *CircularList[T]) Size() int {
	// Implementar
	return -1
}

// IsEmpty evalúa si la lista está vacía.
//
// Uso:
//
//	empty := list.IsEmpty() // Verifica si la lista está vacía.
//
// Retorna:
//   - `true` si la lista está vacía; `false` en caso contrario.
func (l *CircularList[T]) IsEmpty() bool {
	return l.size == 0
}

// Clear elimina todos los elementos de la lista.
//
// Uso:
//
//	list.Clear() // Elimina todos los elementos de la lista.
func (l *CircularList[T]) Clear() {
	// Implementar
}

// Prepend agrega un nuevo nodo al principio de la lista.
//
// Uso:
//
//	list.Prepend(10) // Agrega el valor 10 al principio de la lista.
//
// Parámetros:
//   - `data`: el valor a agregar al principio de la lista.
func (l *CircularList[T]) Prepend(data T) {
	// Implementar
}

// Append agrega un nuevo nodo al final de la lista.
//
// Uso:
//
//	list.Append(10) // Agrega el valor 10 al final de la lista.
//
// Parámetros:
//   - `data`: el valor a agregar al final de la lista.
func (l *CircularList[T]) Append(data T) {
	// Implementar
}

// Find busca un nodo en la lista.
//
// Uso:
//
//	node := list.Find(10) // Busca el valor 10 en la lista.
//
// Parámetros:
//   - `data`: el valor a buscar en la lista.
//
// Retorna:
//   - el nodo que contiene el valor buscado; `nil` si el valor no se encuentra en la lista.
func (l *CircularList[T]) Find(data T) *DoubleLinkedNode[T] {
	// Implementar
	return nil
}

// RemoveFirst elimina el primer nodo de la lista.
//
// Uso:
//
//	list.RemoveFirst() // Elimina el primer nodo de la lista.
func (l *CircularList[T]) RemoveFirst() {
	// Implementar
}

// RemoveLast elimina el último nodo de la lista.
//
// Uso:
//
//	list.RemoveLast() // Elimina el último nodo de la lista.
func (l *CircularList[T]) RemoveLast() {
	// Implementar
}

// Remove elimina un la primera aparición de un dato en la lista.
//
// Uso:
//
//	list.Remove(10) // Elimina la primera aparición del dato 10 en la lista.
//
// Parámetros:
//   - `data`: el dato a eliminar de la lista.
func (l *CircularList[T]) Remove(data T) {
	// Implementar
}

// String devuelve una representación en cadena de la lista.
//
// Uso:
//
//	fmt.Println(list) // Imprime la representación en cadena de la lista.
//
// Retorna:
//   - una representación en cadena de la lista.
func (l *CircularList[T]) String() string {
	// Implementar
	return ""
}
