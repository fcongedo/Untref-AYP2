package ejercicios

import (
	"untref-ayp2/guia-pilas-colas/queue"
	"untref-ayp2/guia-pilas-colas/stack"
)

// Escribir una función que que reciba una cadena de caracteres
// y devuelva la cadena invertida. Analizar el orden.
func InvertirCadena(cadena string) string {
	pila := stack.NewStack[rune]()
	for _, letra := range cadena {
		pila.Push(letra)
	}
	invertida := make([]rune, 0, len(cadena))
	for !pila.IsEmpty() {
		letra, _ := pila.Pop()
		invertida = append(invertida, letra)
	}
	return string(invertida)
}

// Escribir una función que verifique si una palabra es palíndromo
// (es decir que una cadena es igual a su inversa. Por ejemplo:
//
//	las cadenas "1456541" y "145541" son palíndromos). Analizar el orden.
func Palindromo(cadena string) bool {
	pila := stack.NewStack[rune]()

	for _, letra_original := range cadena {
		pila.Push(letra_original)
	}

	for _, letra_original := range cadena {
		letra_pila, _ := pila.Pop()
		if letra_original != letra_pila {
			return false
		}
	}

	return true
}

// Escribir una función que evalúe si una cadena de paréntesis, corchetes y
// llaves está bien balanceada y devuelve `true` si está bien balanceada y
// `false` si está mal balanceada. Por ejemplo: `[()]{}{[()()]()}` debe
// devolver `true`, mientras que `[(])` debe devolver `false`. Analizar el orden.
func Balanceada(cadena string) bool {
	pila := stack.NewStack[string]()
	for _, caracter := range cadena {
		switch caracter {
		case '(', '[', '{':
			pila.Push(string(caracter))
		case ')':
			if pila.IsEmpty() {
				return false
			}
			tope, _ := pila.Pop()
			if tope != "(" {
				return false
			}
		case '}':
			if pila.IsEmpty() {
				return false
			}
			tope, _ := pila.Pop()
			if tope != "{" {
				return false
			}
		case ']':
			if pila.IsEmpty() {
				return false
			}
			tope, _ := pila.Pop()
			if tope != "[" {
				return false
			}
		}

	}
	return pila.IsEmpty()
}

// Escribir una función, tal que, dadas dos colas, construya una cola con el resultado
// de poner una a continuación de la otra. Por ejemplo: si `q1 = [1,2,3]` (con 1 al
// frente y 3 al final) y `q2 = [5,7]`, el resultado es `[1,2,3,5,7]` (con 1 al
// frente y 7 al final). Analizar el orden.
func UnirColas[T any](q1, q2 *queue.Queue[T]) *queue.Queue[T] {
	nueva_cola := queue.NewQueue[T]()

	for !q1.IsEmpty() {
		elemento, _ := q1.Dequeue()
		nueva_cola.Enqueue(elemento)
	}

	for !q2.IsEmpty() {
		elemento, _ := q2.Dequeue()
		nueva_cola.Enqueue(elemento)
	}

	return nueva_cola
}
