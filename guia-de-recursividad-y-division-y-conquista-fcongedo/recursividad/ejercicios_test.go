package recursividad

import (
	"testing"
	"untref-ayp2/guia-recursividad-division-conquista/queue"

	"github.com/stretchr/testify/assert"
)

func TestSuma(t *testing.T) {
	assert.Equal(t, 1, Suma(1))
	assert.Equal(t, 21, Suma(6))
	assert.Equal(t, 55, Suma(10))
}

func TestFactorial(t *testing.T) {
	assert.Equal(t, 2, Factorial(2))
	assert.Equal(t, 24, Factorial(4))
	assert.Equal(t, 120, Factorial(5))
}

func TestCantidadDeUnos(t *testing.T) {
	assert.Equal(t, 0, CantidadDeUnos(0))
	assert.Equal(t, 1, CantidadDeUnos(4))
	assert.Equal(t, 2, CantidadDeUnos(5))
	assert.Equal(t, 3, CantidadDeUnos(42))
}

func TestInvertir(t *testing.T) {
	assert.Equal(t, "neuqueN", Invertir("Neuquen"))
	assert.Equal(t, "!odnuM ,aloH", Invertir("Hola, Mundo!"))
}

func TestInvertirCola(t *testing.T) {
	q := queue.New(1, 2, 3, 4, 5)
	InvertirCola(q)
	assert.Equal(t, []int{5, 4, 3, 2, 1}, q.Values())

	q2 := queue.New[int]()
	InvertirCola(q2)
	assert.Equal(t, []int{}, q2.Values())

	q3 := queue.New(1)
	InvertirCola(q3)
	assert.Equal(t, []int{1}, q3.Values())
}

func TestMCD(t *testing.T) {
	assert.Equal(t, 12, MCD(12, 60))
	assert.Equal(t, 12, MCD(60, 12))
	assert.Equal(t, 1, MCD(7, 13))
	assert.Equal(t, 1, MCD(13, 7))
	assert.Equal(t, 2, MCD(2, 4))
	assert.Equal(t, 2, MCD(4, 2))
	assert.Equal(t, 3, MCD(3, 9))
	assert.Equal(t, 3, MCD(9, 3))
}

func TestMultiplicar(t *testing.T) {
	assert.Equal(t, 16, Multiplicar(2, 8))
	assert.Equal(t, 16, Multiplicar(8, 2))
	assert.Equal(t, 0, Multiplicar(0, 8))
	assert.Equal(t, 0, Multiplicar(8, 0))
}

func TestDivisionEntera(t *testing.T) {
	cociente, resto := DivisionEntera(8, 2)
	assert.Equal(t, 4, cociente)
	assert.Equal(t, 0, resto)
}

func TestSumaArray(t *testing.T) {
	assert.Equal(t, 0, SumaArray([]int{}))
	assert.Equal(t, 1, SumaArray([]int{1}))
	assert.Equal(t, 3, SumaArray([]int{1, 2}))
	assert.Equal(t, 6, SumaArray([]int{1, 2, 3}))
}

func TestDecimalABinario(t *testing.T) {
	assert.Equal(t, "101010", DecimalABinario(42))
	assert.Equal(t, "10000000000000000000000000000000", DecimalABinario(2147483648))
	assert.Equal(t, "0", DecimalABinario(0))
	assert.Equal(t, "1", DecimalABinario(1))
	assert.Equal(t, "1100", DecimalABinario(12))
}

func TestEsPotencia(t *testing.T) {
	assert.True(t, EsPotencia(8, 2))
	assert.True(t, EsPotencia(1, 2))
	assert.False(t, EsPotencia(7, 2))
}

func TestFibonacci(t *testing.T) {
	assert.Equal(t, 5, Fibonacci(5))
	assert.Equal(t, 8, Fibonacci(6))
	assert.Equal(t, 13, Fibonacci(7))
	assert.Equal(t, 21, Fibonacci(8))
}

func TestPascal(t *testing.T) {
	assert.Equal(t, 1, Pascal(0, 0))
	assert.Equal(t, 1, Pascal(1, 0))
	assert.Equal(t, 1, Pascal(1, 1))
	assert.Equal(t, 2, Pascal(2, 1))
	assert.Equal(t, 3, Pascal(3, 2))
	assert.Equal(t, 6, Pascal(4, 2))
	assert.Equal(t, 10, Pascal(5, 2))
}
