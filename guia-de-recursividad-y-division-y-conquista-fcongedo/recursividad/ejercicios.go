package recursividad

import "untref-ayp2/guia-recursividad-division-conquista/queue"

func Suma(n int) int {
	if n == 1 {
		return 1
	}
	return n + Suma(n-1)
}

func Factorial(n int) int {
	if n == 1 || n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}

func CantidadDeUnos(n int) int {
	if n == 0 {
		return 0
	}
	return (n % 2) + CantidadDeUnos(n/2)
}

func Invertir(cadena string) string {
	if len(cadena) <= 1 {
		return cadena
	}
	return Invertir(cadena[1:]) + string(cadena[0])

}

func InvertirCola[T any](q *queue.Queue[T]) {
	if q.IsEmpty() {
		return
	}
	elem := q.Dequeue()

	InvertirCola(q)

	q.Enqueue(elem)
}

func MCD(a, b int) int {
	if b == 0 {
		return a
	}
	return MCD(b, a%b)
}

func Multiplicar(a, b int) int {
	if b == 0 {
		return 0
	}
	return a + Multiplicar(a, b-1)
}

func DivisionEntera(dividendo, divisor int) (cociente, resto int) {
	// Implementar
	return -1, -1
}

func SumaArray(v []int) int {
	// Implementar
	return -1
}

func DecimalABinario(n int) string {
	// Implementar
	return ""
}

func EsPotencia(n, b int) bool {
	// Implementar
	return false
}

func Fibonacci(n int) int {
	// Implementar
	return -1
}

func Pascal(n, k int) int {
	// Implementar
	return -1
}
