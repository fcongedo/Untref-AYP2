// Package stack proporciona una implementación de una pila genérica sobre un arreglo dinámico.
// Expone la estructura y funciones básicas para manipular una pila.
package stack

// Stack proporciona una pila cuyos elementos son de un tipo genérico.
// La implementación se basa en un arreglo dinámico.
type Stack[T any] struct {
	// Implementar
}

// NewStack crea una nueva pila vacía.
//
// Uso:
//
//	s := stack.New[int]() // Crea una pila de enteros.
func NewStack[T any]() *Stack[T] {
	// Implementar
	return nil
}

// Push agrega un elemento a la pila.
//
// Uso:
//
//	s.Push(10) // Agrega el entero 10 a la pila.
func (s *Stack[T]) Push(x T) {
	// Implementar
}

// Pop remueve y retorna el elemento en el tope de la pila.
// Si la pila está vacía, retorna un error.
//
// Uso:
//
//	if x, err := s.Pop(); err != nil {
//		fmt.Println(err)
//	} else {
//		fmt.Println(x)
//	}
func (s *Stack[T]) Pop() (T, error) {
	// Implementar
	var x T
	return x, nil
}

// Top retorna el elemento en el tope de la pila.
// Si la pila está vacía, retorna un error.
//
// Uso:
//
//	if x, err := s.Top(); err != nil {
//		fmt.Println(err)
//	} else {
//		fmt.Println(x)
//	}
func (s *Stack[T]) Top() (T, error) {
	// Implementar
	var x T
	return x, nil
}

// IsEmpty retorna true si la pila está vacía.
//
// Uso:
//
//	if s.IsEmpty() {
//		fmt.Println("La pila está vacía")
//	}
func (s *Stack[T]) IsEmpty() bool {
	// Implementar
	return false
}
