package sortedlist

import (
	"cmp"
	"untref-ayp2/guia-listas/list"
)

// Lista ordenada de forma ascendente
type SortedList[T cmp.Ordered] struct {
	lista *list.DoubleLinkedList[T]
}

func NewSortedList[T cmp.Ordered]() *SortedList[T] {
	return &SortedList[T]{list.NewDoubleLinkedList[T]()}
}

func (s *SortedList[T]) Insert(data T) {
	// Implementar
}

func (s *SortedList[T]) RemoveFirst() {
	s.lista.RemoveFirst()
}

func (s *SortedList[T]) RemoveLast() {
	s.lista.RemoveLast()
}

func (s *SortedList[T]) Remove(data T) {
	s.lista.Remove(data)
}

func (s *SortedList[T]) IsEmpty() bool {
	return s.lista.IsEmpty()
}

func (s *SortedList[T]) Size() int {
	return s.lista.Size()
}

func (s *SortedList[T]) Head() T {
	return s.lista.Head().Data()
}

func (s *SortedList[T]) Tail() T {
	return s.lista.Tail().Data()
}

func (s *SortedList[T]) Find(data T) *list.DoubleLinkedNode[T] {
	return s.lista.Find(data)
}

func (s *SortedList[T]) Clear() {
	s.lista.Clear()
}
