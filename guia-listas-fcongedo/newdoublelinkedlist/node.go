package newdoublelinkedlist

type Node[T comparable] struct {
	data T
	next *Node[T]
	prev *Node[T]
}

func NewNode[T comparable](dato T) *Node[T] {
	// Implementar
	return nil
}

func (n *Node[T]) Data() T {
	// Implementar
	var x T
	return x
}

func (n *Node[T]) Next() *Node[T] {
	// Implementar
	return nil
}

func (n *Node[T]) Prev() *Node[T] {
	// Implementar
	return nil
}

func (n *Node[T]) SetData(dato T) {
	// Implementar
}

func (n *Node[T]) SetNext(sig *Node[T]) {
	// Implementar
}

func (n *Node[T]) SetPrev(prev *Node[T]) {
	// Implementar
}

func (n *Node[T]) InsertAfter(dato T) {
	// Implementar
}

func (n *Node[T]) InsertBefore(dato T) {
	// Implementar
}

func (n *Node[T]) Remove() {
	// Implementar
}

func (n *Node[T]) RemoveNext() {
	// Implementar
}

func (n *Node[T]) RemovePrev() {
	// Implementar
}
